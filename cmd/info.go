package cmd

import (
	"cider/internal/commands/info"
	"context"

	"github.com/urfave/cli/v3"
)

func newInfo() *cli.Command {
	return &cli.Command{
		Name:    "info",
		Usage:   "Display information about a range",
		Aliases: []string{"f"},
		Action: func(_ context.Context, command *cli.Command) error {
			args := command.Args().Slice()

			handler := info.New()

			return handler.Handle(args)
		},
		UsageText: "info [range]",
	}
}
