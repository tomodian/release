package task

import (
	"fmt"

	"release/cmd/commandkey"
	"release/cmd/flagkey"
	"release/files"
	"release/parser"
	"release/utils"

	"github.com/urfave/cli/v2"
)

func Unreleased(workdir string) *cli.Command {
	return &cli.Command{
		Name:    commandkey.Unreleased,
		Usage:   fmt.Sprintf("List all changes for %s", parser.Unreleased),
		Aliases: []string{"u"},
		Flags: []cli.Flag{
			dirFlag(workdir),
			ignoreEmptyFlag(workdir),
		},
		Action: func(c *cli.Context) error {

			for _, p := range files.Glob(c.String(flagkey.Directory)) {
				doc, err := files.Read(p)

				if err != nil {
					return err
				}

				outs, err := parser.Show(doc, parser.Unreleased)

				if err != nil {
					return err
				}

				if c.Bool(flagkey.IgnoreEmpty) && len(outs) == 0 {
					continue
				}

				fmt.Println(utils.Pretty(files.Rel(p)))
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
