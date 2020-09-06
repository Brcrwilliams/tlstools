package main

import (
	"certreader/cli"
	"os"
)

func main() {
	err := cli.NewCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
