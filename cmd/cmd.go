package cmd

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/urfave/cli/v3"
)

func Execute(stdout, stderr io.Writer, args []string) {
	cmd := &cli.Command{
		Name:  "cider",
		Usage: "cli tool to help with common IP related tasks",
		Commands: []*cli.Command{
			newRanges(stdout),
			newIn(stdout),
			newSubnet(stdout),
			newInfo(stdout),
			newVersion(stdout),
		},
	}

	// if the first arg is a number, treat it as the "ranges" command
	if len(args) == 2 {
		maybeNumber := args[1]
		if _, err := strconv.Atoi(maybeNumber); err == nil {
			args[1] = "ranges"
			args = append(args, maybeNumber)
		}
	}

	if err := cmd.Run(context.Background(), args); err != nil {
		fmt.Fprint(stderr, err)
	}
}
