package main

import (
	"cider/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Stdout, os.Stderr, os.Args)
}
