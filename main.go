package main

import (
	"demyst-todo/cmd"
	"demyst-todo/log"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	log.Init()
	app := cli.NewApp()
	app.Name = "demyst-todo"
	app.Usage = "A command line tool that consumes a todo and outputs the title and status."

	app.Commands = []*cli.Command{
		cmd.CMDStatus(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}
}
