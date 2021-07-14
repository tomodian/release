package files

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmatcuk/doublestar"
)

var (
	excludes = []string{
		"__pycache__",
		"_compareTemp",
		"_notes",
		".cache",
		".dynamodb",
		".eggs",
		".grunt",
		".idea",
		".ipynb_checkpoints",
		".mypy_cache",
		".next",
		".npm",
		".parcel-cache",
		".phpunit.result.cache",
		".prof",
		".sass-cache",
		".scrapy",
		".terraform",
		".vagrant",
		".vendor",
		".vs",
		".vscode",
		".vuepress",
		"bower_components",
		"build",
		"coverage",
		"dist",
		"DocProject",
		"htmlcov",
		"jspm_packages",
		"node_modules",
		"vendor",
		"x64",
		"x86",
	}
)

func ignore(path string) bool {
	slash := filepath.ToSlash(path)

	for _, e := range excludes {
		chunks := strings.Split(slash, "/")

		for _, c := range chunks {
			if c == e {
				return true
			}
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

// Rel returns relative path.
func Rel(path string) string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	rel, err := filepath.Rel(wd, path)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("## %s", rel)
}
