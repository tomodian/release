package main

import (
	"fmt"
	"log"
	"os"

	"release/files"
	"release/parser"
	"release/utils"

	"github.com/manifoldco/promptui"
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

	verFlagRequired := &cli.StringFlag{
		Name:     verFlag.Name,
		Usage:    "target `VERSION`",
		Aliases:  []string{"v"},
		Required: true,
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

						doc, err := files.Read(p)

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
					verFlagRequired,
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range utils.Glob(c.String("dir")) {
						fmt.Println(chalk.Magenta.Color(p))
						fmt.Println("")

						doc, err := files.Read(p)

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
					verFlagRequired,
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					v, err := parser.Version(c.String("version"))

					if err != nil {
						return err
					}

					targets := utils.Glob(c.String("dir"))

					if len(targets) == 0 {
						fmt.Println("(nothing found)")
						return nil
					}

					fmt.Println(chalk.Magenta.Color("Targets"))

					for _, p := range utils.Glob(c.String("dir")) {
						fmt.Println(p)
					}

					prompt := promptui.Select{
						Label: chalk.Magenta.Color(fmt.Sprintf("Do you really want to bump these CHANGELOGs to version %s?", v)),
						Items: []string{"no", "yes"},
					}

					_, picked, err := prompt.Run()

					if err != nil {
						return err
					}

					if picked != "yes" {
						fmt.Println("Cancelled")
						return nil
					}

					for _, p := range utils.Glob(c.String("dir")) {
						fmt.Printf("%s --> ", chalk.Magenta.Color(p))

						doc, err := files.Read(p)

						if err != nil {
							return err
						}

						body, err := parser.To(doc, c.String("version"))

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
