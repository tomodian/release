package cmd

import (
	"fmt"
	"log"
	"os"

	"release/cmd/commandkey"
	"release/cmd/flagkey"
	"release/cmd/header"
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

	anyFlag := &cli.BoolFlag{
		Name:    flagkey.Any,
		Value:   false,
		Usage:   "ignore semantic versioning and grab anything inside [string]",
		Aliases: []string{"a"},
	}

	dirFlag := &cli.StringFlag{
		Name:    flagkey.Directory,
		Value:   wd,
		Usage:   "target `DIR`",
		Aliases: []string{"d"},
	}

	ignoreEmptyFlag := &cli.BoolFlag{
		Name:    flagkey.IgnoreEmpty,
		Value:   false,
		Usage:   "ignore output for changelog without changes",
		Aliases: []string{"i"},
	}

	verFlag := &cli.StringFlag{
		Name:    flagkey.Version,
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

	typeFlag := &cli.StringFlag{
		Name:    flagkey.Type,
		Usage:   "Semver `TYPE`: X.Y.Z refers to {major}.{minor}.{patch} or {release}.{feature}.{hotfix}",
		Aliases: []string{"t"},
	}

	app := &cli.App{
		Name:  "release",
		Usage: "Manage changelog for your release process 🚀",
		Commands: []*cli.Command{
			{
				Name:    commandkey.Targets,
				Usage:   "List all CHANGELOG.md files",
				Aliases: []string{"target", "t"},
				Flags: []cli.Flag{
					dirFlag,
				},
				Action: func(c *cli.Context) error {
					fmt.Println(utils.Pretty(header.Target))

					for _, p := range files.Glob(c.String(flagkey.Directory)) {
						fmt.Println(files.Rel(p))
					}

					return nil
				},
			},
			{
				Name:    commandkey.Latest,
				Usage:   "Show the latest released version in current directory",
				Aliases: []string{"l"},
				Flags: []cli.Flag{
					anyFlag,
					dirFlag,
				},
				Action: func(c *cli.Context) error {

					doc, err := files.Read(fmt.Sprintf("%s/CHANGELOG.md", wd))

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

					fmt.Println(got)

					return nil
				},
			},
			{
				Name:    commandkey.Unreleased,
				Usage:   fmt.Sprintf("List all changes for %s", parser.Unreleased),
				Aliases: []string{"u"},
				Flags: []cli.Flag{
					dirFlag,
					ignoreEmptyFlag,
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
			},
			{
				Name:  commandkey.To,
				Usage: "Bump all [Unreleased] sections to given version",
				Flags: []cli.Flag{
					verFlagRequired,
					dirFlag,
					&cli.BoolFlag{
						Name:    flagkey.Force,
						Usage:   "Force without prompt, Mainly for CI environment",
						Aliases: []string{"f"},
					},
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
			},
			{
				Name:    commandkey.Next,
				Usage:   "Suggest next version by checking CHANGELOGs recursively",
				Aliases: []string{"n"},
				Flags: []cli.Flag{
					dirFlag,
					typeFlag,
				},
				Action: func(c *cli.Context) error {

					latests := map[string]parser.SemanticVersion{}

					// Construct a map of versions.
					for _, p := range files.Glob(c.String(flagkey.Directory)) {
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

					tflag := c.String(flagkey.Type)

					if tflag == "" {
						fmt.Println("")
						fmt.Println("Latest released version:", chalk.Magenta.Color(ver.String()))
						fmt.Println("")
						fmt.Println("Suggestions for next release:")
						fmt.Println("   - Major / Release -->", chalk.Magenta.Color((ver.Increment(parser.MajorVersion).String())))
						fmt.Println("   - Minor / Feature -->", chalk.Magenta.Color(ver.Increment(parser.MinorVersion).String()))
						fmt.Println("   - Patch / Hotfix  -->", chalk.Magenta.Color(ver.Increment(parser.PatchVersion).String()))
						fmt.Println("")

						return nil
					}

					vtype, err := parser.AliasedVersion(tflag)

					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					switch vtype {

					case parser.MajorVersion:
						fmt.Print(ver.Increment(parser.MajorVersion).String())
						return nil

					case parser.MinorVersion:
						fmt.Print(ver.Increment(parser.MinorVersion).String())
						return nil

					case parser.PatchVersion:
						fmt.Print(ver.Increment(parser.PatchVersion).String())
						return nil
					}

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
