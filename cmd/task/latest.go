package task

import (
	"fmt"

	"github.com/tomodian/release/cmd/commandkey"
	"github.com/tomodian/release/cmd/flagkey"
	"github.com/tomodian/release/cmd/task/flag"
	"github.com/tomodian/release/files"
	"github.com/tomodian/release/parser"

	"github.com/urfave/cli/v2"
)

func Latest(workdir string) *cli.Command {
	return &cli.Command{
		Name:    commandkey.Latest,
		Usage:   "Show the latest released version in current directory",
		Aliases: []string{"l"},
		Flags: []cli.Flag{
			flag.Any(workdir),
			flag.Dir(workdir),
			flag.Newline(workdir, true),
		},
		Action: func(c *cli.Context) error {

			doc, err := files.Read(fmt.Sprintf("%s/CHANGELOG.md", workdir))

			if err != nil {
				return err
			}

			if c.Bool(flagkey.Any) {
				got, err := parser.LatestAny(doc)

				if err != nil {
					return err
				}

				fmt.Println(got)

				return nil
			}

			got, err := parser.Latest(doc)

			if err != nil {
				return err
			}

			if c.Bool(flagkey.Newline) {
				fmt.Println(got)
				return nil
			}

			fmt.Print(got)

			return nil
		},
	}
}
