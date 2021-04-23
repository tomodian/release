package cmd

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

// Run it.
func Run(args []string) error {
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

	ignoreEmptyFlag := &cli.BoolFlag{
		Name:    flagIgnoreEmpty,
		Value:   false,
		Usage:   "ignore output for changelog without changes",
		Aliases: []string{"i"},
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

	versionFlag := &cli.StringFlag{
		Name:        flagType,
		Usage:       "Semver `TYPE`: X.Y.Z refers to {major}.{minor}.{patch}",
		Aliases:     []string{"t"},
		DefaultText: parser.MinorVersion.String(),
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
					fmt.Println(utils.Pretty(headingTarget))

					for _, p := range files.Glob(c.String(flagDirectory)) {
						fmt.Println(files.Rel(p))
					}

					return nil
				},
			},
			{
				Name:    cmdLatest,
				Usage:   "Show the latest released version in current directory",
				Aliases: []string{"l"},
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					doc, err := files.Read(fmt.Sprintf("%s/CHANGELOG.md", wd))

					if err != nil {
						return err
					}

					got, err := parser.Latest(doc)

					if err != nil {
						return err
					}

					fmt.Println(got)

					return nil
				},
			},
			{
				Name:    cmdNext,
				Usage:   fmt.Sprintf("List all changes for %s", parser.Unreleased),
				Aliases: []string{"n"},
				Flags: []cli.Flag{
					dirFlag,
					ignoreEmptyFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range files.Glob(c.String(flagDirectory)) {
						doc, err := files.Read(p)

						if err != nil {
							return err
						}

						outs, err := parser.Show(doc, parser.Unreleased)

						if err != nil {
							return err
						}

						if c.Bool(flagIgnoreEmpty) && len(outs) == 0 {
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
			},
			{
				Name:    "show",
				Usage:   "Show changes of given version",
				Aliases: []string{"s"},
				Flags: []cli.Flag{
					verFlagRequired,
					dirFlag,
					ignoreEmptyFlag,
				},
				Action: func(c *cli.Context) error {

					for _, p := range files.Glob(c.String(flagDirectory)) {
						doc, err := files.Read(p)

						if err != nil {
							return err
						}

						outs, err := parser.Show(doc, c.String(flagVersion))

						if err != nil {
							return err
						}

						if c.Bool(flagIgnoreEmpty) && len(outs) == 0 {
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
						fmt.Println(files.Rel(p))
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
						fmt.Printf("%s --> ", files.Rel(p))

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
			{
				Name:    cmdVersion,
				Usage:   "Suggest next version by checking CHANGELOGs recursively",
				Aliases: []string{"v"},
				Flags: []cli.Flag{
					dirFlag,
					versionFlag,
				},
				Action: func(c *cli.Context) error {

					latests := map[string]parser.SemanticVersion{}

					// Construct a map of versions.
					for _, p := range files.Glob(c.String(flagDirectory)) {
						doc, err := files.Read(p)

						if err != nil {
							return err
						}

						lat, err := parser.Latest(doc)

						if err != nil {
							continue
						}

						if _, exists := latests[lat]; exists {
							continue
						}

						v, err := parser.NewSemanticVersion(lat)

						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}

						latests[lat] = *v
					}

					vers := []parser.SemanticVersion{}

					for k := range latests {
						vers = append(vers, latests[k])
					}

					vers = parser.SortVersions(vers)
					ver := vers[len(vers)-1]

					fmt.Println("")
					fmt.Println("Latest released version:", chalk.Magenta.Color(ver.String()))
					fmt.Println("")
					fmt.Println("Suggestions for next release:")
					fmt.Println("   - Major / Release -->", chalk.Magenta.Color((ver.Increment(parser.MajorVersion).String())))
					fmt.Println("   - Minor / Feature -->", chalk.Magenta.Color(ver.Increment(parser.MinorVersion).String()))
					fmt.Println("   - Patch / Hotfix  -->", chalk.Magenta.Color(ver.Increment(parser.PatchVersion).String()))
					fmt.Println("")

					return nil
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
		return err
	}

	return nil
}
