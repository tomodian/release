package parser_test

import (
	"testing"

	"release/parser"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	falsyVersions = []string{
		"",
		"1",
		"1.0",
		"1.0.X",
		"1.0.0.0",
		"1a.2.3",
		"1.2a.3",
		"1.2.3a",
		"1 2 3",
		"🍎",
	}
)

func TestNewSemanticVersion(t *testing.T) {
	{
		// Success cases.
		type pattern struct {
			expected parser.SemanticVersion
			sample   string
		}

		pats := []pattern{
			{
				expected: parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				sample:   "0.0.0",
			},
			{
				expected: parser.SemanticVersion{Major: 1, Minor: 1, Patch: 1},
				sample:   "1.1.1",
			},
			{
				expected: parser.SemanticVersion{Major: 123, Minor: 456, Patch: 789},
				sample:   "123.456.789",
			},
		}

		for _, p := range pats {
			got, err := parser.NewSemanticVersion(p.sample)

			require.Nilf(t, err, spew.Sdump(p))
			require.NotNilf(t, got, spew.Sdump(p))

			assert.Equalf(t, p.expected.Major, got.Major, spew.Sdump(p, got))
			assert.Equalf(t, p.expected.Minor, got.Minor, spew.Sdump(p, got))
			assert.Equalf(t, p.expected.Patch, got.Patch, spew.Sdump(p, got))
		}
	}

	{
		// Fail cases.
		for _, p := range falsyVersions {
			_, err := parser.NewSemanticVersion(p)

			require.NotNilf(t, err, "%s", p)
		}
	}
}

func TestNewSemanticVersionIsEqual(t *testing.T) {
	type pattern struct {
		a *parser.SemanticVersion
		b *parser.SemanticVersion
	}

	{
		// Truthy cases.
		pats := []pattern{
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 1, Minor: 2, Patch: 3},
				b: &parser.SemanticVersion{Major: 1, Minor: 2, Patch: 3},
			},
		}

		for _, p := range pats {
			assert.Truef(t, p.a.IsEqual(p.b), spew.Sdump(p))
		}
	}

	{
		// Falsy cases.
		pats := []pattern{
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				b: nil,
			},
			{
				a: &parser.SemanticVersion{Major: 4, Minor: 5, Patch: 6},
				b: &parser.SemanticVersion{Major: 1, Minor: 2, Patch: 3},
			},
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 1},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 1, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 1, Minor: 0, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
		}

		for _, p := range pats {
			assert.Falsef(t, p.a.IsEqual(p.b), spew.Sdump(p))
		}
	}
}

func TestNewSemanticVersionIsGreater(t *testing.T) {
	type pattern struct {
		a *parser.SemanticVersion
		b *parser.SemanticVersion
	}

	{
		// Truthy cases.
		pats := []pattern{
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 1},
			},
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 1, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				b: &parser.SemanticVersion{Major: 1, Minor: 0, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 1, Minor: 2, Patch: 3},
				b: &parser.SemanticVersion{Major: 4, Minor: 5, Patch: 6},
			},
		}

		for _, p := range pats {
			assert.Truef(t, p.a.IsGreater(p.b), spew.Sdump(p))
		}
	}

	{
		// Falsy cases.
		pats := []pattern{
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
				b: nil,
			},
			{
				a: &parser.SemanticVersion{Major: 4, Minor: 5, Patch: 6},
				b: &parser.SemanticVersion{Major: 1, Minor: 2, Patch: 3},
			},
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 1},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 0, Minor: 1, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
			{
				a: &parser.SemanticVersion{Major: 1, Minor: 0, Patch: 0},
				b: &parser.SemanticVersion{Major: 0, Minor: 0, Patch: 0},
			},
		}

		for _, p := range pats {
			assert.Falsef(t, p.a.IsGreater(p.b), spew.Sdump(p))
		}
	}
}

func TestCastVersion(t *testing.T) {
	{
		// Success cases.
		type pattern struct {
			expected int
			sample   string
		}

		pats := []pattern{
			{
				expected: 0,
				sample:   "0",
			},
			{
				expected: 1,
				sample:   "1",
			},
		}

		for _, p := range pats {
			got, err := parser.CastVersion("foo", p.sample)

			require.NoErrorf(t, err, spew.Sdump(p))
			assert.Equalf(t, p.expected, got, spew.Sdump(p))
		}
	}

	{
		// Fail cases.
		pats := []string{
			"",
			"-2",
			"broken",
		}

		for _, p := range pats {
			got, err := parser.CastVersion("foo", p)

			require.Errorf(t, err, spew.Sdump(p))
			assert.Equalf(t, -1, got, spew.Sdump(p))
		}
	}
}
