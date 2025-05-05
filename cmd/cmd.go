package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v3"
)

func Execute(stdin io.Reader, stdout, stderr io.Writer) {
	cmd := &cli.Command{
		Name:  "cider",
		Usage: "cli tool to help with common IP related tasks",
		Commands: []*cli.Command{
			newRanges(),
			newIn(),
			newSubnet(),
			newInfo(),
			newVersion(),
		},
		Writer:    stdout,
		ErrWriter: stderr,
		Reader:    stdin,
	}

	args := os.Args

	// if the first arg is a number, treat it as the "ranges" command
	if len(args) == 2 {
		maybeNumber := args[1]
		if _, err := strconv.Atoi(maybeNumber); err == nil {
			args[1] = "ranges"
			args = append(args, maybeNumber)
		}
	}

	if err := cmd.Run(context.Background(), args); err != nil {
		log.Fatal(err)
	}
}
