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
				Usage:   "display cidr ranges",
				Aliases: []string{"r"},
				Action: func(*cli.Context) error {
					handler := ranges.New()

					return handler.Handle()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
