package cmd

import (
	"cider/internal/commands/in"
	"context"
	"io"

	"github.com/urfave/cli/v3"
)

func newIn(stdout io.Writer) *cli.Command {
	return &cli.Command{
		Name:    "in",
		Usage:   "Determine if an ip or range falls within one or more ranges",
		Aliases: []string{"i"},
		Action: func(_ context.Context, command *cli.Command) error {
			args := command.Args().Slice()

			handler := in.New(stdout)

			return handler.Handle(args)
		},
		UsageText: "in [ip or range] [range1] [optional range2] [optional rangeN]",
	}
}
