package cmd

import (
	"demyst-todo/handlers"

	"github.com/urfave/cli/v2"
)

func CMDStatus() *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "Returns the status of the TODO",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "limit",
				Aliases: []string{"l"},
				Value:   20,
				Usage:   "The number of todos to fetch",
			},
			&cli.StringFlag{
				Name:    "pattern",
				Aliases: []string{"p"},
				Usage:   "The pattern to filter the todos. Available Pattern: even, odd, all",
				Value:   "all",
			},
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Value:   "api",
				Usage:   "The input source of the todos. Available values: api, file",
			},
			&cli.StringFlag{
				Name:    "location",
				Aliases: []string{"loc"},
				Value:   "https://jsonplaceholder.typicode.com/todos",
				Usage:   "The location to fetch the todos from. It can either be file location or API url depending on the input",
			},
		},
		Action: handlers.StatusHandler,
	}
}
