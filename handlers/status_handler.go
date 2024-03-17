package handlers

import (
	"demyst-todo/input"
	"demyst-todo/types"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

func parseFlag(ctx *cli.Context) (types.Config, error) {
	limit := ctx.Int("limit")
	if limit <= 0 {
		return types.Config{}, fmt.Errorf("limit cannot be 0 or negative")
	}

	pattern := ctx.String("pattern")
	if pattern != "even" && pattern != "odd" && pattern != "all" {
		return types.Config{}, fmt.Errorf("pattern must be even, odd or all")
	}

	input := ctx.String("input")
	if input != "api" && input != "file" {
		return types.Config{}, fmt.Errorf("input must be api or file")
	}

	location := ctx.String("location")
	if input == "api" && location == "" {
		return types.Config{}, fmt.Errorf("location cannot be empty")
	}

	return types.Config{
		Limit:    limit,
		Pattern:  pattern,
		Location: location,
		Input:    input,
	}, nil
}

func StatusHandler(ctx *cli.Context) error {
	cfg, err := parseFlag(ctx)
	if err != nil {
		return err
	}

	var inp input.Input
	if cfg.Input == "api" {
		inp = &input.API{URL: cfg.Location}
	} else {
		inp = &input.File{Location: cfg.Location}
	}

	todos, err := inp.Fetch(cfg.Limit, cfg.Pattern)
	if err != nil {
		return err
	}

	for _, todo := range todos {
		todoJSON, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		fmt.Println(string(todoJSON))
	}

	return nil
}
