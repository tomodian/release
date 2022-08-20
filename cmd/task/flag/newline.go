package flag

import (
	"release/cmd/flagkey"

	"github.com/urfave/cli/v2"
)

func Newline(workdir string, val bool) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:  flagkey.Newline,
		Usage: "add newline at the end of certain output",
		Value: val,
	}
}
