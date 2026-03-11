package task

import (
	"fmt"

	"github.com/tomodian/release/cmd/commandkey"
	"github.com/tomodian/release/cmd/flagkey"
	"github.com/tomodian/release/cmd/header"
	"github.com/tomodian/release/cmd/task/flag"
	"github.com/tomodian/release/files"
	"github.com/tomodian/release/utils"

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
