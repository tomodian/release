package parser_test

import (
	"fmt"
	"os"
	"testing"

	"release/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	wd, err := os.Getwd()

	require.NoError(t, err)

	{
		// Success cases.
		pats := []string{
			// Refer to the top-level directory of this project.
			fmt.Sprintf("%s/../", wd),
		}

		for _, p := range pats {
			got, err := parser.NewRepository(p)

			require.NoErrorf(t, err, p)
			assert.NotNilf(t, got, p)
		}
	}

	{
		// Fail cases.
		pats := []string{
			"",
			"./",
		}

		for _, p := range pats {
			got, err := parser.NewRepository(p)

			require.Errorf(t, err, p)
			assert.Nilf(t, got, p)
		}
	}
}

func TestRepositoryTagNames(t *testing.T) {

	wd, err := os.Getwd()

	require.NoError(t, err)

	// Refer to the top-level directory of this project.
	repo, err := parser.NewRepository(fmt.Sprintf("%s/../", wd))

	require.NoError(t, err)

	{
		// Success case.
		got, err := repo.TagNames()

		require.NoError(t, err)
		assert.NotEmpty(t, got)

		// Testing against the real git tag result.
		assert.Contains(t, got, parser.SemanticVersion{Major: 0, Minor: 1, Patch: 7})
	}

	{
		// Fail cases.
		// TODO
	}
}
