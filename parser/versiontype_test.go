package parser_test

import (
	"testing"

	"release/parser"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionTypeString(t *testing.T) {
	assert.Equal(t, string(parser.MajorVersion), parser.MajorVersion.String())
}

func TestAliasedVersion(t *testing.T) {
	{
		// Success cases, MajorVersion
		pats := []string{
			parser.MajorVersion.String(),
			"release",
		}

		for _, p := range pats {
			got, err := parser.AliasedVersion(p)

			require.NoErrorf(t, err, spew.Sdump(p))
			assert.Equal(t, parser.MajorVersion, got, spew.Sdump(p, got))
		}
	}

	{
		// Success cases, MinorVersion
		pats := []string{
			parser.MinorVersion.String(),
			"feature",
		}

		for _, p := range pats {
			got, err := parser.AliasedVersion(p)

			require.NoErrorf(t, err, spew.Sdump(p))
			assert.Equal(t, parser.MinorVersion, got, spew.Sdump(p, got))
		}
	}

	{
		// Success cases, PatchVersion
		pats := []string{
			parser.PatchVersion.String(),
			"hotfix",
		}

		for _, p := range pats {
			got, err := parser.AliasedVersion(p)

			require.NoErrorf(t, err, spew.Sdump(p))
			assert.Equal(t, parser.PatchVersion, got, spew.Sdump(p, got))
		}
	}

	{
		// Fail cases
		pats := []string{
			"",
			"foo",
		}

		for _, p := range pats {
			got, err := parser.AliasedVersion(p)

			require.Errorf(t, err, spew.Sdump(p))
			assert.Empty(t, got, spew.Sdump(p, got))
		}
	}
}
