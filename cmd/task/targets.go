package task

import (
	"fmt"

	"release/cmd/commandkey"
	"release/cmd/flagkey"
	"release/cmd/header"
	"release/cmd/task/flag"
	"release/files"
	"release/utils"

	"github.com/urfave/cli/v2"
)

// Targets returns all CHANGELOG.md files.
func Targets(workdir string) *cli.Command {
	return &cli.Command{
		Name:    commandkey.Targets,
		Usage:   "List all CHANGELOG.md files",
		Aliases: []string{"target", "t"},
		Flags: []cli.Flag{
			flag.Dir(workdir),
		},
		Action: func(c *cli.Context) error {
			fmt.Println(utils.Pretty(header.Target))

			for _, p := range files.Glob(c.String(flagkey.Directory)) {
				fmt.Println(files.Rel(p))
			}

			return nil
		},
	}
}
