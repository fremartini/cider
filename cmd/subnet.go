package cmd

import (
	"cider/internal/commands/subnet"
	"context"
	"io"

	"github.com/urfave/cli/v3"
)

func newSubnet(stdout io.Writer) *cli.Command {
	return &cli.Command{
		Name:    "subnet",
		Usage:   "Split a range into multiple smaller ranges",
		Aliases: []string{"s"},
		Action: func(_ context.Context, command *cli.Command) error {
			args := command.Args().Slice()

			handler := subnet.New(stdout)

			return handler.Handle(args)
		},
		UsageText: "subnet [range] [size1] [optional size2] [optional sizeN]",
	}
}
