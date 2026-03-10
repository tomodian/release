package main

import (
	"os"

	"github.com/tomodian/release/cmd"
)

func main() {
	if err := cmd.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
