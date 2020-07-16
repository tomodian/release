package utils

import (
	"fmt"
	"sort"

	"github.com/bmatcuk/doublestar"
)

// Glob seeks for all CHANGELOG.md in given directory.
func Glob(d string) []string {
	p := fmt.Sprintf("%s/**/CHANGELOG.md", d)

	paths, err := doublestar.Glob(p)

	if err != nil {
		panic("malformed path pattern")
	}

	sort.Slice(paths, func(i, j int) bool {
		return paths[i] < paths[j]
	})

	return paths
}
