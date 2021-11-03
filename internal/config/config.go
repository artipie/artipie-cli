package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"gopkg.in/yaml.v3"
)

// ArtiCtlConfig is a config for artictl.
type ArtiCtlConfig struct {
	CurrentContext string                 `yaml:"currentContext"`
	Contexts       map[string]*CtlContext `yaml:"contexts"`
}

// CtlContext represents credentials for access to the Artipie server.
type CtlContext struct {
	Auth     Auth   `yaml:"-"`
	Endpoint string `yaml:"endpoint"`
}

type errConfigNotFound struct{ paths []string }

func (e *errConfigNotFound) Error() string {
	return fmt.Sprintf("file not found, paths: %v",
		strings.Join(e.paths, ", "))
}

type errConfigFieldEmpty struct{ field string }

func (e *errConfigFieldEmpty) Error() string {
	return fmt.Sprintf("config field %s is empty", e.field)
}

// FromFiles parses config from first existent of the specifed files.
func (c *ArtiCtlConfig) FromFiles(paths ...string) error {
	for _, p := range paths {
		if p == "" {
			continue
		}
		p = os.ExpandEnv(p)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			continue
		}
		return c.parse(p)
	}
	return &errConfigNotFound{paths}
}

func (c *ArtiCtlConfig) parse(fileName string) error {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}
	return c.validate()
}

func (c *ArtiCtlConfig) validate() error {
	var me *multierror.Error
	if c.CurrentContext == "" {
		me = multierror.Append(me, &errConfigFieldEmpty{"currentContext"})
	}
	if len(c.Contexts) == 0 {
		me = multierror.Append(me, &errConfigFieldEmpty{"contexts"})
	}
	for k, v := range c.Contexts {
		err := v.validate()
		if err != nil {
			err := fmt.Errorf("invalid %v context spec: %w", k, err)
			me = multierror.Append(me, err)
		}
	}
	return me.ErrorOrNil()
}

func (c *CtlContext) validate() error {
	var me *multierror.Error
	if c.Endpoint == "" {
		me = multierror.Append(me, &errConfigFieldEmpty{"endpoint"})
	}
	err := c.Auth.validate()
	if err != nil {
		me = multierror.Append(me, err)
	}
	return me.ErrorOrNil()
}

// ContextFromInput sets CtlContext fields with specified values.
func (c *CtlContext) ContextFromInput(user, password, token, endpoint string) error {
	if endpoint != "" {
		c.Endpoint = endpoint
	}
	if token != "" && user == "" && password == "" {
		c.Auth = AuthToken{token}
		return nil
	}
	if user != "" && password != "" && token == "" {
		c.Auth = AuthBasic{user, password}
		return nil
	}
	if user == "" && password == "" && token == "" {
		return nil
	}
	return &errInvalidAuth{}
}
