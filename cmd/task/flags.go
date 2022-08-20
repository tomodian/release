package task

import (
	"release/cmd/flagkey"
	"release/parser"

	"github.com/urfave/cli/v2"
)

func anyFlag(workdir string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    flagkey.Any,
		Value:   false,
		Usage:   "ignore semantic versioning and grab anything inside [string]",
		Aliases: []string{"a"},
	}
}

func dirFlag(workdir string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    flagkey.Directory,
		Value:   workdir,
		Usage:   "target `DIR`",
		Aliases: []string{"d"},
	}
}

func ignoreEmptyFlag(workdir string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    flagkey.IgnoreEmpty,
		Value:   false,
		Usage:   "ignore output for changelog without changes",
		Aliases: []string{"i"},
	}
}

func typeFlag(workdir string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    flagkey.Type,
		Usage:   "Semver `TYPE`: X.Y.Z refers to {major}.{minor}.{patch} or {release}.{feature}.{hotfix}",
		Aliases: []string{"t"},
	}
}

func verFlag(workdir string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    flagkey.Version,
		Usage:   "target `VERSION`",
		Value:   parser.Unreleased,
		Aliases: []string{"v"},
	}
}

func verFlagRequired(workdir string) *cli.StringFlag {
	v := verFlag(workdir)

	return &cli.StringFlag{
		Name:     v.Name,
		Usage:    v.Usage,
		Aliases:  v.Aliases,
		Required: true,
	}
}
