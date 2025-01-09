package main

import (
	"cider/internal/commands/in"
	"cider/internal/commands/ranges"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cider",
		Usage: "CIDR cli tool",
		Commands: []*cli.Command{
			{
				Name:    "ranges",
				Usage:   "Displays all CIDR ranges",
				Aliases: []string{"r"},
				Action: func(c *cli.Context) error {
					arg := c.Args().First()

					handler := ranges.New()

					return handler.Handle(arg)
				},
				UsageText: "ranges [optional range]",
			},
			{
				Name:    "in",
				Usage:   "Determines if an ip falls within a range",
				Aliases: []string{"i"},
				Action: func(c *cli.Context) error {
					args := c.Args().Slice()

					handler := in.New()

					return handler.Handle(args)
				},
				UsageText: "in [ip] [range1] [optional range2] [optional rangeN]",
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

	if err := app.Run(args); err != nil {
		log.Fatal(err)
	}
}
