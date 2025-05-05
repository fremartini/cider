package main

import (
	"cider/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Stdin, os.Stdout, os.Stderr)
}
