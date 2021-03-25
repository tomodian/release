package utils

import (
	"strings"

	"github.com/ttacon/chalk"
)

// List of common lines.
const (
	EmptyLine = "(empty)"
)

// Pretty color on terminal.
func Pretty(in string) string {
	switch {
	case in == EmptyLine:
		return chalk.Dim.TextStyle(EmptyLine)

	case strings.HasPrefix(in, "# "):
		return chalk.Yellow.Color(in)

	case strings.HasPrefix(in, "## "):
		return chalk.Blue.Color(in)

	case strings.HasPrefix(in, "### "):
		return chalk.Cyan.Color(in)
	}

	return in
}
