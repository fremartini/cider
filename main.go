package main

import (
	"cider/internal/commands/in"
	"cider/internal/commands/info"
	"cider/internal/commands/ranges"
	"cider/internal/commands/subnet"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v3"
)

var version string

func main() {
	cmd := &cli.Command{
		Name:  "cider",
		Usage: "CIDR cli tool",
		Commands: []*cli.Command{
			{
				Name:    "ranges",
				Usage:   "display all CIDR ranges",
				Aliases: []string{"r"},
				Action: func(_ context.Context, command *cli.Command) error {
					arg := command.Args().First()

					handler := ranges.New()

					return handler.Handle(arg)
				},
				UsageText: "ranges [optional range]",
			},
			{
				Name:    "in",
				Usage:   "determine if an ip falls within a range",
				Aliases: []string{"i"},
				Action: func(_ context.Context, command *cli.Command) error {
					args := command.Args().Slice()

					handler := in.New()

					return handler.Handle(args)
				},
				UsageText: "in [ip] [range1] [optional range2] [optional rangeN]",
			},
			{
				Name:    "subnet",
				Usage:   "split a range into multiple smaller ranges",
				Aliases: []string{"s"},
				Action: func(_ context.Context, command *cli.Command) error {
					args := command.Args().Slice()

					handler := subnet.New()

					return handler.Handle(args)
				},
				UsageText: "subnet [range] [size1] [optional size2] [optional sizeN]",
			},
			{
				Name:    "info",
				Usage:   "display information about an ip or range",
				Aliases: []string{"f"},
				Action: func(_ context.Context, command *cli.Command) error {
					args := command.Args().Slice()

					handler := info.New()

					return handler.Handle(args)
				},
				UsageText: "info [range or ip]",
			},
			{
				Name:        "version",
				Aliases:     []string{"v"},
				Description: "show version",
				Action: func(_ context.Context, command *cli.Command) error {
					fmt.Println(version)

					return nil
				},
			},
		},
	}

	args := os.Args

	// if the first arg is a number, treat it as the "ranges" command
	if len(args) == 2 {
		maybeNumber := args[1]
		if _, err := strconv.Atoi(maybeNumber); err == nil {
			args[1] = "ranges"
			args = append(args, maybeNumber)
		}
	}

	if err := cmd.Run(context.Background(), args); err != nil {
		log.Fatal(err)
	}
}
