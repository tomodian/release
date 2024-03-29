package files_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/tomodian/release/files"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGlob(t *testing.T) {

	pwd, err := os.Getwd()

	require.Nil(t, err)

	type pattern struct {
		path  string
		count int
	}

	{
		// Fail cases, panic.
		pats := []pattern{
			{
				path: "[-]",
			},
		}

		for _, p := range pats {
			assert.Panics(t, func() {
				files.Glob(p.path)
			})
		}
	}

	{
		// Fail case, ensure vendor directories are not included.
		path := fmt.Sprintf("%s/test/vendors", pwd)

		assert.Empty(t, files.Glob(path))
	}

	{
		// Success cases.
		pats := []pattern{
			{
				path:  fmt.Sprintf("%s/test", pwd),
				count: 5,
			},
			{
				path:  fmt.Sprintf("%s/test/some", pwd),
				count: 3,
			},
			{
				path:  fmt.Sprintf("%s/test/some/nested", pwd),
				count: 2,
			},
			{
				path:  fmt.Sprintf("%s/test/some/nested/directory", pwd),
				count: 1,
			},
			{
				path:  fmt.Sprintf("%s/test/some/nested/directory/NotExistent", pwd),
				count: 0,
			},
		}

		for _, p := range pats {
			got := files.Glob(p.path)

			require.Nilf(t, err, "%s", p)
			assert.Equalf(t, p.count, len(got), "%s", p.path)
		}
	}
}

func TestRel(t *testing.T) {
	type pattern struct {
		path string
	}

	{
		// Success cases.
		pwd, err := os.Getwd()

		require.Nil(t, err)

		pats := []pattern{
			{
				path: fmt.Sprintf("%s/test", pwd),
			},
			{
				path: fmt.Sprintf("%s/test/some", pwd),
			},
			{
				path: fmt.Sprintf("%s/test/some/nested", pwd),
			},
			{
				path: fmt.Sprintf("%s/test/some/nested/directory", pwd),
			},
			{
				path: fmt.Sprintf("%s/test/some/nested/directory/NotExistent", pwd),
			},
		}

		for _, p := range pats {
			assert.NotPanics(t, func() {
				files.Rel(p.path)
			})
		}
	}
}
