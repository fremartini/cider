package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/urfave/cli/v3"
)

var version string

func newVersion(stdout io.Writer) *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show version",
		Action: func(_ context.Context, command *cli.Command) error {
			fmt.Fprintln(stdout, version)

			return nil
		},
	}
}
