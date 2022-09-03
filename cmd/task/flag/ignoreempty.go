package flag

import (
	"github.com/tomodian/release/cmd/flagkey"

	"github.com/urfave/cli/v2"
)

func IgnoreEmpty(workdir string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    flagkey.IgnoreEmpty,
		Value:   false,
		Usage:   "ignore output for changelog without changes",
		Aliases: []string{"i"},
	}
}
