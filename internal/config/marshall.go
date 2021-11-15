package config

import (
	"gopkg.in/yaml.v3"
)

func (c *CtlContext) UnmarshalYAML(value *yaml.Node) error {
	type C CtlContext
	type temp struct {
		Auth yaml.Node `yaml:"auth"`
		*C   `yaml:",inline"`
	}
	obj := &temp{C: (*C)(c)}
	if err := value.Decode(obj); err != nil {
		return err
	}
	content := obj.Auth.Content
	if len(content) == 2 && content[0].Value == "token" {
		c.Auth = new(AuthToken)
	} else if len(content) == 4 && content[0].Value == "username" && content[2].Value == "password" {
		c.Auth = new(AuthBasic)
	} else {
		return ErrInvalidAuth
	}
	err := obj.Auth.Decode(c.Auth)
	return err
}

func (c *CtlContext) MarshalYAML() (interface{}, error) {
	type C CtlContext
	type temp struct {
		Auth interface{} `yaml:"auth"`
		*C   `yaml:",inline"`
	}
	t := &temp{c.Auth, (*C)(c)}
	return t, nil
}
