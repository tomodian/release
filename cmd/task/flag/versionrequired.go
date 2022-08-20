package flag

import (
	"github.com/urfave/cli/v2"
)

func VersionRequired(workdir string) *cli.StringFlag {
	v := Version(workdir)

	return &cli.StringFlag{
		Name:     v.Name,
		Usage:    v.Usage,
		Aliases:  v.Aliases,
		Required: true,
	}
}
