package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

const (
	configLocal = "$HOME/.config/artictl/config.yaml"
	configSys   = "/etc/artictl/config.yaml"
)

var app = cli.App{
	Name:        "artictl",
	Description: "CLI tool for managing Artipie server",
	Action:      run(),
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "context",
			Usage:   "Select non-default context",
			Aliases: []string{"ctx"},
		},
		&cli.StringFlag{
			Name:    "endpoint",
			Usage:   "Endpoint URI",
			Aliases: []string{"e"},
		},
		&cli.StringFlag{
			Name:    "auth-user",
			Usage:   "User for authentification",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    "auth-password",
			Usage:   "Password for authentification",
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:    "auth-token",
			Usage:   "Token for authentification",
			Aliases: []string{"t"},
		},
	},
	Commands: []*cli.Command{
		{
			Name:  "get",
			Usage: "Display one or many resources",
			Subcommands: []*cli.Command{
				{
					Name:  "repo",
					Usage: "List repositories",
				},
				{
					Name:  "users",
					Usage: "List users",
				},
				{
					Name:  "perm",
					Usage: "List permissions",
				},
				{
					Name:  "storage",
					Usage: "List storages",
				},
			},
		},
		{
			Name:  "update",
			Usage: "Create or update a resource",
			Subcommands: []*cli.Command{
				{
					Name:  "repo",
					Usage: "Create repository",
				},
				{
					Name:  "users",
					Usage: "Create user",
				},
				{
					Name:  "perm",
					Usage: "Create permission",
				},
				{
					Name:  "storage",
					Usage: "Create storage",
				},
			},
		},
		{
			Name:  "delete",
			Usage: "Delete a resource",
			Subcommands: []*cli.Command{
				{
					Name:  "repo",
					Usage: "Delete repository",
				},
				{
					Name:  "users",
					Usage: "Delete user",
				},
				{
					Name:  "perm",
					Usage: "Delete permission",
				},
				{
					Name:  "storage",
					Usage: "Delete storage",
				},
			},
		},
	},
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		c, err := resolveContext(ctx)
		if err != nil {
			err := fmt.Errorf("Failed to parse config: %w", err)
			return err
		}
		err = yaml.NewEncoder(os.Stdout).Encode(c)
		if err != nil {
			return err
		}
		return nil
	}
}
