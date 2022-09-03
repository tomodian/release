package flag

import (
	"github.com/tomodian/release/cmd/flagkey"

	"github.com/urfave/cli/v2"
)

func Type(workdir string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    flagkey.Type,
		Usage:   "Semver `TYPE`: X.Y.Z refers to {major}.{minor}.{patch} or {release}.{feature}.{hotfix}",
		Aliases: []string{"t"},
	}
}
