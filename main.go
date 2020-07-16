package main

import (
	"fmt"
	"log"
	"os"

	"release/parser"
	"release/utils"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func main() {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dirFlag := &cli.StringFlag{
		Name:    "dir",
		Value:   wd,
		Usage:   "target `DIR`",
		Aliases: []string{"d"},
	}

	verFlag := &cli.StringFlag{
		Name:    "version",
		Value:   parser.Unreleased,
		Aliases: []string{"v"},
	}

	app := &cli.App{
		Name:  "release",
		Usage: "Manage changelog for your release process🚀",
		Commands: []*cli.Command{
			{
				Action: func(c *cli.Context) error {
					fmt.Println("boom! I say!")
					return nil
				},
			},
			{
				Name:    "targets",
				Usage:   "List all CHANGELOG.md files",
				Aliases: []string{"t"},
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {
					fmt.Println(chalk.Magenta.Color("Targets:"))

					for _, p := range utils.Glob(c.String("dir")) {
						fmt.Println(p)
					}

					return nil
				},
			},
			{
				Name:    "next",
				Usage:   fmt.Sprintf("List all changes for %s", parser.Unreleased),
				Aliases: []string{"n"},
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range utils.Glob(c.String("dir")) {
						fmt.Println(chalk.Magenta.Color(p))
						fmt.Println("")

						doc, err := utils.ReadFile(p)

						if err != nil {
							return err
						}

						outs, err := parser.Show(doc, parser.Unreleased)

						if err != nil {
							return err
						}

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
			},
			{
				Name:    "show",
				Usage:   "Show changes of given version",
				Aliases: []string{"for"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    verFlag.Name,
						Aliases: []string{"v"},
						// Version must be explicitly passed.
						Required: true,
					},
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range utils.Glob(c.String("dir")) {
						fmt.Println(chalk.Magenta.Color(p))
						fmt.Println("")

						doc, err := utils.ReadFile(p)

						if err != nil {
							return err
						}

						outs, err := parser.Show(doc, c.String("version"))

						if err != nil {
							return err
						}

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
			},
			{
				Name:  "to",
				Usage: "Bump all [Unreleased] sections to given version",
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {
					fmt.Println("new task template: ", c.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}
}
