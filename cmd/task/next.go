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
			flag.Newline(workdir, false),
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

			// Print all possible versions when user did not specify the specific type.
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

			var out = ""

			switch vtype {

			case parser.MajorVersion:
				out = ver.Increment(parser.MajorVersion).String()

			case parser.MinorVersion:
				out = ver.Increment(parser.MinorVersion).String()

			case parser.PatchVersion:
				out = ver.Increment(parser.PatchVersion).String()
			}

			if c.Bool(flagkey.Newline) {
				fmt.Println(out)
				return nil
			}

			// Note for zsh users:
			// - Newline is always present.
			// - STDOUT will be suffixed with percentage sign `%`, which could be muted in zsh configuration.
			//   https://unix.stackexchange.com/questions/167582/why-zsh-ends-a-line-with-a-highlighted-percent-symbol
			fmt.Print(out)

			return nil
		},
	}
}
