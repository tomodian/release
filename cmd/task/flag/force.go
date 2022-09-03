package flag

import (
	"github.com/tomodian/release/cmd/flagkey"

	"github.com/urfave/cli/v2"
)

func Force(workdir string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    flagkey.Force,
		Usage:   "force without prompt, Mainly for CI environment",
		Aliases: []string{"f"},
	}
}
