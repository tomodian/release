package files_test

import (
	"fmt"
	"os"
	"testing"

	"release/files"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGlob(t *testing.T) {
	type pattern struct {
		path  string
		count int
	}

	{
		// Success cases.
		pwd, err := os.Getwd()

		require.Nil(t, err)

		pats := []pattern{
			{
				path:  fmt.Sprintf("%s/test", pwd),
				count: 4,
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
