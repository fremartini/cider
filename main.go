package main

import (
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
				Usage:   "display all CIDR ranges",
				Aliases: []string{"r"},
				Action: func(c *cli.Context) error {
					arg := c.Args().First()

					handler := ranges.New()

					return handler.PrintAllCIDRBlocks(arg)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
