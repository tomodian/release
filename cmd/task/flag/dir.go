package flag

import (
	"github.com/tomodian/release/cmd/flagkey"

	"github.com/urfave/cli/v2"
)

func Dir(workdir string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    flagkey.Directory,
		Value:   workdir,
		Usage:   "target `DIR`",
		Aliases: []string{"d"},
	}
}
