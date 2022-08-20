package task

import (
	"fmt"
	"os"

	"release/cmd/commandkey"
	"release/cmd/flagkey"
	"release/cmd/task/flag"
	"release/files"
	"release/parser"

	"github.com/ttacon/chalk"
	"github.com/urfave/cli/v2"
)

func Next(workdir string) *cli.Command {
	return &cli.Command{
		Name:    commandkey.Next,
		Usage:   "Suggest next version by checking CHANGELOGs recursively",
		Aliases: []string{"n"},
		Flags: []cli.Flag{
			flag.Dir(workdir),
			flag.Type(workdir),
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
	}
}
