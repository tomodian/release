package main

import (
	"os"

	"release/cmd"
)

func main() {
	_ = cmd.Run(os.Args)
}
