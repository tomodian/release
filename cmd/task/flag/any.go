package flag

import (
	"github.com/tomodian/release/cmd/flagkey"

	"github.com/urfave/cli/v2"
)

func Any(workdir string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    flagkey.Any,
		Value:   false,
		Usage:   "ignore semantic versioning and grab anything inside [string]",
		Aliases: []string{"a"},
	}
}
