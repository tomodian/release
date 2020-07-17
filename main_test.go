package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// See GitHub discussion for testing CLI apps.
// https://github.com/urfave/cli/issues/731

func TestRun(t *testing.T) {
	{
		// Success cases.
		pats := []string{
			"",
			cmdTargets, "target", "t",
			cmdNext, "n",
			cmdTo,
		}

		for _, p := range pats {
			args := os.Args[0:1]
			args = append(args, p)

			require.NotPanics(t, func() {
				run(args)
			})
		}
	}
}
