package main

import (
	"cider/internal/commands/in"
	"cider/internal/commands/ranges"
	"log"
	"os"

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

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
