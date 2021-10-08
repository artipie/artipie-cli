package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = cli.App{
	Name: "artictl",
	Description: "CLI tool for managing Artipie server",
	Commands: []*cli.Command{
		{
			Name: "get",
			Usage: "Display one or many resources",
			Subcommands: []*cli.Command{
				{
					Name: "repo",
					Usage: "List repositories",
				},
				{
					Name: "users",
					Usage: "List users",
				},
				{
					Name: "perm",
					Usage: "List permissions",
				},
				{
					Name: "storage",
					Usage: "List storages",
				},
			},
		},
		{
			Name: "update",
			Usage: "Create or update a resource",
			Subcommands: []*cli.Command{
				{
					Name: "repo",
					Usage: "Create repository",
				},
				{
					Name: "users",
					Usage: "Create user",
				},
				{
					Name: "perm",
					Usage: "Create permission",
				},
				{
					Name: "storage",
					Usage: "Create storage",
				},
			},
		},
		{
			Name: "delete",
			Usage: "Delete a resource",
			Subcommands: []*cli.Command{
				{
					Name: "repo",
					Usage: "Delete repository",
				},
				{
					Name: "users",
					Usage: "Delete user",
				},
				{
					Name: "perm",
					Usage: "Delete permission",
				},
				{
					Name: "storage",
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
