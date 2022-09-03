package files_test

import (
	"os"
	"testing"

	"github.com/tomodian/release/files"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {

	type pattern struct {
		path string
		doc  string
	}

	{
		// Fail cases.
		pats := []pattern{
			{
				path: "",
				doc:  "",
			},
			{
				path: "./non-existent",
				doc:  "",
			},
			{
				path: "",
				doc:  "whatever",
			},
		}

		for _, p := range pats {
			err := files.Update(p.path, p.doc)

			require.NotNilf(t, err, "%s", err)
		}
	}

	{
		// Fail case, non-existent path.
		err := files.Update("non-existent", "foo")

		require.Nil(t, err)
	}

	{
		// Success cases.
		pats := []pattern{
			{
				path: "test/sample.md",
				doc:  "world",
			},
		}

		for _, p := range pats {
			info, err := os.Stat(p.path)
			require.Nil(t, err)

			assert.Nil(t, files.Update(p.path, p.doc))

			up, err := os.Stat(p.path)
			require.Nil(t, err)

			assert.Equal(t, info.Mode(), up.Mode())
		}
	}
}
