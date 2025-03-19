package cmd

import (
	"cider/internal/commands/in"
	"context"

	"github.com/urfave/cli/v3"
)

func newIn() *cli.Command {
	return &cli.Command{
		Name:    "in",
		Usage:   "Determine if an ip falls within a range",
		Aliases: []string{"i"},
		Action: func(_ context.Context, command *cli.Command) error {
			args := command.Args().Slice()

			handler := in.New()

			return handler.Handle(args)
		},
		UsageText: "in [ip] [range1] [optional range2] [optional rangeN]",
	}
}
