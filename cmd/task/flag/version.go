package flag

import (
	"release/cmd/flagkey"
	"release/parser"

	"github.com/urfave/cli/v2"
)

func Version(workdir string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    flagkey.Version,
		Usage:   "target `VERSION`",
		Value:   parser.Unreleased,
		Aliases: []string{"v"},
	}
}
