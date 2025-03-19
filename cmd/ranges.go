package cmd

import (
	"cider/internal/commands/ranges"
	"context"

	"github.com/urfave/cli/v3"
)

func newRanges() *cli.Command {
	return &cli.Command{
		Name:    "ranges",
		Usage:   "Display all CIDR ranges",
		Aliases: []string{"r"},
		Action: func(_ context.Context, command *cli.Command) error {
			arg := command.Args().First()

			handler := ranges.New()

			return handler.Handle(arg)
		},
		UsageText: "ranges [optional range]",
	}
}
