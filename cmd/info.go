package cmd

import (
	"cider/internal/commands/info"
	"context"
	"io"

	"github.com/urfave/cli/v3"
)

func newInfo(stdout io.Writer) *cli.Command {
	return &cli.Command{
		Name:    "info",
		Usage:   "Display information about a range",
		Aliases: []string{"f"},
		Action: func(_ context.Context, command *cli.Command) error {
			args := command.Args().Slice()

			handler := info.New(stdout)

			return handler.Handle(args)
		},
		UsageText: "info [range]",
	}
}
