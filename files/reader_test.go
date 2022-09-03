package files_test

import (
	"testing"

	"github.com/tomodian/release/files"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	{
		// Fail cases.
		pats := []string{
			"",
			"./non-existent",
			"whatever",
		}

		for _, p := range pats {
			_, err := files.Read(p)

			require.NotNilf(t, err, "%s", err)
		}
	}

	{
		// Success cases.
		pats := []string{
			"test/sample.md",
		}

		for _, p := range pats {
			doc, err := files.Read(p)

			require.Nil(t, err)
			assert.NotEmpty(t, doc)
		}
	}
}
