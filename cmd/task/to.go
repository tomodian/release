package task

import (
	"fmt"

	"github.com/tomodian/release/cmd/commandkey"
	"github.com/tomodian/release/cmd/flagkey"
	"github.com/tomodian/release/cmd/header"
	"github.com/tomodian/release/cmd/task/flag"
	"github.com/tomodian/release/files"
	"github.com/tomodian/release/parser"

	"github.com/manifoldco/promptui"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func To(workdir string) *cli.Command {
	return &cli.Command{
		Name:  commandkey.To,
		Usage: "Bump all [Unreleased] sections to given version",
		Flags: []cli.Flag{
			flag.VersionRequired(workdir),
			flag.Dir(workdir),
			flag.Force(workdir),
		},
		Action: func(c *cli.Context) error {

			v, err := parser.Version(c.String(flagkey.Version))

			if err != nil {
				return err
			}

			targets := files.Glob(c.String(flagkey.Directory))

			if len(targets) == 0 {
				fmt.Println("(nothing found)")
				return nil
			}

			fmt.Println(chalk.Magenta.Color(header.Target))

			for _, p := range files.Glob(c.String(flagkey.Directory)) {
				fmt.Println(files.Rel(p))
			}

			if !c.Bool(flagkey.Force) {
				agreed := "yes"

				prompt := promptui.Prompt{
					Label: chalk.Magenta.Color(fmt.Sprintf("Enter `%s` to update all CHANGELOGs to version %s", agreed, v)),
				}

				picked, err := prompt.Run()

				if err != nil {
					return err
				}

				if picked != agreed {
					fmt.Println("Cancelled")
					return nil
				}
			}

			fmt.Println("")

			for _, p := range files.Glob(c.String(flagkey.Directory)) {
				fmt.Printf("%s --> ", files.Rel(p))

				doc, err := files.Read(p)

				if err != nil {
					return err
				}

				body, err := parser.To(doc, c.String(flagkey.Version))

				if err != nil {
					return err
				}

				if len(body) == 0 {
					fmt.Println("skipped")
					continue
				}

				if err := files.Update(p, body); err != nil {
					fmt.Println("❌")
					continue
				}

				fmt.Println("✅")
			}

			fmt.Println("Done👍")

			return nil
		},
	}
}
