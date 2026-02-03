package main

import (
	"fmt"
	"os"

	"alidns/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:], cli.NewDefaultDeps(os.Stdout, os.Stderr)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
