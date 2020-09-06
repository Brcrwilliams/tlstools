package main

import (
	"os"

	"github.com/brcrwilliams/tlstools/cmd/x509meta/cli"
)

func main() {
	err := cli.NewCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
