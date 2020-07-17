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

const (
	// Headings
	headingTarget = "Targets:"

	// Subcommands
	cmdTargets = "targets"
	cmdNext    = "next"
	cmdTo      = "to"

	// Flag keys
	flagDirectory = "dir"
	flagVersion   = "version"
	flagForce     = "force"
)

func run(args []string) {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dirFlag := &cli.StringFlag{
		Name:    flagDirectory,
		Value:   wd,
		Usage:   "target `DIR`",
		Aliases: []string{"d"},
	}

	verFlag := &cli.StringFlag{
		Name:    flagVersion,
		Usage:   "target `VERSION`",
		Value:   parser.Unreleased,
		Aliases: []string{"v"},
	}

	verFlagRequired := &cli.StringFlag{
		Name:     verFlag.Name,
		Usage:    verFlag.Usage,
		Aliases:  verFlag.Aliases,
		Required: true,
	}

	app := &cli.App{
		Name:  "release",
		Usage: "Manage changelog for your release process 🚀",
		Commands: []*cli.Command{
			{
				Name:    cmdTargets,
				Usage:   "List all CHANGELOG.md files",
				Aliases: []string{"target", "t"},
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {
					fmt.Println(chalk.Magenta.Color(headingTarget))

					for _, p := range files.Glob(c.String(flagDirectory)) {
						fmt.Println(p)
					}

					return nil
				},
			},
			{
				Name:    cmdNext,
				Usage:   fmt.Sprintf("List all changes for %s", parser.Unreleased),
				Aliases: []string{"n"},
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range files.Glob(c.String(flagDirectory)) {
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
				Aliases: []string{"s"},
				Flags: []cli.Flag{
					verFlagRequired,
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range files.Glob(c.String(flagDirectory)) {
						fmt.Println(chalk.Magenta.Color(p))
						fmt.Println("")

						doc, err := files.Read(p)

						if err != nil {
							return err
						}

						outs, err := parser.Show(doc, c.String(flagVersion))

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
				Name:  cmdTo,
				Usage: "Bump all [Unreleased] sections to given version",
				Flags: []cli.Flag{
					verFlagRequired,
					dirFlag,
					&cli.BoolFlag{
						Name:    flagForce,
						Usage:   "Force without prompt, Mainly for CI environment",
						Aliases: []string{"f"},
					},
				},
				Action: func(c *cli.Context) error {

					v, err := parser.Version(c.String(flagVersion))

					if err != nil {
						return err
					}

					targets := files.Glob(c.String(flagDirectory))

					if len(targets) == 0 {
						fmt.Println("(nothing found)")
						return nil
					}

					fmt.Println(chalk.Magenta.Color(headingTarget))

					for _, p := range files.Glob(c.String(flagDirectory)) {
						fmt.Println(p)
					}

					if !c.Bool(flagForce) {
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

					for _, p := range files.Glob(c.String(flagDirectory)) {
						fmt.Printf("%s --> ", p)

						doc, err := files.Read(p)

						if err != nil {
							return err
						}

						body, err := parser.To(doc, c.String(flagVersion))

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
			},
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}
}

func main() {
	run(os.Args)
}
