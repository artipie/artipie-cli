package main

import (
	"fmt"

	"github.com/artipie/artipie-cli/internal/config"
	"github.com/urfave/cli/v2"
)

func resolveContext(ctx *cli.Context) (*config.CtlContext, error) {
	endpoint := ctx.String("endpoint")
	user := ctx.String("auth-user")
	password := ctx.String("auth-password")
	token := ctx.String("auth-token")
	currentContext := ctx.String("context")
	ctlctx := &config.CtlContext{}

	if !fullContextProvided(user, password, token, endpoint) {
		c, err := contextFromConfig(currentContext)
		if err != nil {
			return nil, err
		}
		ctlctx = c
	}
	err := ctlctx.ContextFromInput(user, password, token, endpoint)
	if err != nil {
		return nil, err
	}
	return ctlctx, nil
}

func contextFromConfig(currentContext string) (*config.CtlContext, error) {
	cfg := config.ArtiCtlConfig{}
	err := cfg.FromFiles(configLocal, configSys)
	if err != nil {
		return nil, err
	}
	if currentContext == "" {
		currentContext = cfg.CurrentContext
	}
	c, ok := cfg.Contexts[currentContext]
	if !ok {
		return nil, fmt.Errorf("context %v not found", currentContext)
	}
	return c, nil
}

func fullContextProvided(user, password, token, endpoint string) bool {
	if endpoint == "" {
		return false
	}
	return token != "" || user != "" && password != ""
}
