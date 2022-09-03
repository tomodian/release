package task

import (
	"fmt"

	"github.com/tomodian/release/cmd/flagkey"
	"github.com/tomodian/release/cmd/task/flag"
	"github.com/tomodian/release/files"
	"github.com/tomodian/release/parser"
	"github.com/tomodian/release/utils"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func Show(workdir string) *cli.Command {
	return &cli.Command{
		Name:    "show",
		Usage:   "Show changes of given version",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			flag.VersionRequired(workdir),
			flag.Dir(workdir),
			flag.IgnoreEmpty(workdir),
		},
		Action: func(c *cli.Context) error {

			for _, p := range files.Glob(c.String(flagkey.Directory)) {
				doc, err := files.Read(p)

				if err != nil {
					return err
				}

				outs, err := parser.Show(doc, c.String(flagkey.Version))

				if err != nil {
					return err
				}

				if c.Bool(flagkey.IgnoreEmpty) && len(outs) == 0 {
					continue
				}

				fmt.Println(chalk.Magenta.Color(files.Rel(p)))
				fmt.Println("")

				if len(outs) == 0 {
					fmt.Println(utils.Pretty(utils.EmptyLine))
					fmt.Println("")
					continue
				}

				for _, o := range outs {
					fmt.Println(utils.Pretty(o))
				}

				fmt.Println("")
			}

			return nil
		},
	}
}
