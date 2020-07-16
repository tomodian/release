package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"release/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	{
		// Success cases.
		pats := []string{
			"0.1.0",
			"0.1.0-beta",
			"100.200.300",
		}

		for _, p := range pats {
			got, err := parser.Version(p)

			require.Nilf(t, err, "%s", p)
			assert.Equal(t, fmt.Sprintf("[%s]", p), got)
		}
	}

	{
		// Fail cases.
		pats := []string{
			"",
			"1",
			"1.0",
			"1.0.X",
			"1.0.0.0",
			"🍎",
		}

		for _, p := range pats {
			_, err := parser.Version(p)

			require.NotNilf(t, err, "%s", p)
		}
	}
}

func TestTo(t *testing.T) {
	type pattern struct {
		version string
		doc     string
	}

	{
		// Success cases.
		pats := []pattern{
			{
				version: "0.1.0",
				doc: strings.Join([]string{
					"# Hello",
					"",
					"## [Unreleased]",
					"",
					"### Added",
					"",
				}, "\n"),
			},
		}

		for _, p := range pats {
			got, err := parser.To(p.doc, p.version)
			require.Nil(t, err)

			v, err := parser.Version(p.version)
			require.Nil(t, err)

			assert.Equal(t, 1, strings.Count(got, v))
		}
	}

	{
		// Fail cases.
		pats := []pattern{
			{
				version: "",
				doc:     strings.Join([]string{}, "\n"),
			},
			{
				version: "0.1.BROKEN",
				doc: strings.Join([]string{
					"# Hello",
					"",
					"## [Unreleased]",
				}, "\n"),
			},
			{
				version: "0.1.0",
				doc: strings.Join([]string{
					"# Hello",
				}, "\n"),
			},
			{
				version: "0.1.BROKEN",
				doc: strings.Join([]string{
					"# Hello",
					"",
					"## [Unreleased]",
					"## [Unreleased]", // Duplicated
				}, "\n"),
			},
		}

		for _, p := range pats {
			_, err := parser.To(p.doc, p.version)

			require.NotNil(t, err)
		}
	}
}

func TestShow(t *testing.T) {
	type pattern struct {
		version string
		doc     string
		count   int
	}

	{
		// Success cases.
		pats := []pattern{
			{
				version: "0.1.0",
				doc: strings.Join([]string{
					"# Hello",
				}, "\n"),
				count: 0,
			},
			{
				version: "0.1.0",
				doc: strings.Join([]string{
					"# Hello",
					"",
					"## [Unreleased]",
					"",
					"## [0.1.0]",
					"",
					"### Added",
					"- foo",
				}, "\n"),
				count: 2,
			},
			{
				version: "1.2.3",
				doc: strings.Join([]string{
					"# Hello",
					"",
					"## [Unreleased]",
					"",
					"## [1.2.3] - 2020/07/16",
					"",
					"### Added",
					"- foo",
					"",
					"### Deleted",
					"- foo",
					"",
					"## [0.1.0]",
					"",
				}, "\n"),
				count: 5,
			},
		}

		for _, p := range pats {
			gots, err := parser.Show(p.doc, p.version)

			require.Nil(t, err)
			assert.Equalf(t, p.count, len(gots), "%#v", gots)
		}
	}

	{
		// Fail cases.
		pats := []pattern{
			{
				version: "",
				doc:     strings.Join([]string{}, "\n"),
			},
		}

		for _, p := range pats {
			_, err := parser.Show(p.doc, p.version)

			require.NotNil(t, err)
		}
	}
}
