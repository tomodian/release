package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalIgnore(t *testing.T) {
	{
		// Truthy patterns.
		pats := append([]string{
			"/node_modules",
			"/nested/.cache/dir",
			`c:\Windows\node_modules`,
		}, excludes...)

		for _, p := range pats {
			got := ignore(p)

			assert.Truef(t, got, "%s", p)
		}
	}

	{
		// Falsy patterns.
		pats := []string{
			"/some/path/for/",
		}

		for _, p := range pats {
			got := ignore(p)

			assert.Falsef(t, got, "%s", p)
		}
	}
}
