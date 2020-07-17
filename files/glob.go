package files

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar"
)

var (
	excludes = []string{
		".cache",
		".vagrant",
		".vendor",
		"build",
		"coverage",
		"node_modules",
	}
)

func ignore(path string) bool {
	slash := filepath.ToSlash(path)

	for _, e := range excludes {
		if strings.Contains(slash, e) {
			return true
		}
	}

	return false
}

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

	outs := []string{}

	// Exclude vendor directories.
	for _, p := range paths {

		if ignore(p) {
			continue
		}

		outs = append(outs, p)
	}

	return outs
}
