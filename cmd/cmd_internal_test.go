package cmd

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
				// TODO: handle errors
				_ = Run(args)
			})
		}
	}

	{
		// Fail case for `next` task.
		pats := [][]string{
			{"next", "--dir", "non-existent"},
		}

		for _, p := range pats {
			os.Args = p

			require.NotNilf(t, Run(os.Args), "#%v", p)
		}
	}

	{
		// Fail case for `show` task.
		pats := [][]string{
			{"show", "--dir", "non-existent"},
			{"show", "--dir", "non-existent", "--version", "x.y.z"},
		}

		for _, p := range pats {
			os.Args = p

			require.NotNilf(t, Run(os.Args), "#%v", p)
		}
	}

	{
		// Fail case for `to` task.
		pats := [][]string{
			{"to", "--dir", "non-existent"},
			{"to", "--dir", "non-existent", "--version", "x.x.x"},
		}

		for _, p := range pats {
			os.Args = p

			require.NotNilf(t, Run(os.Args), "#%v", p)
		}
	}

	{
		// Fail case for `version` task.
		pats := [][]string{
			{"version", "--dir", "non-existent"},
		}

		for _, p := range pats {
			os.Args = p

			require.NotNilf(t, Run(os.Args), "#%v", p)
		}
	}
}
