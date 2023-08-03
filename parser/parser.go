package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/blang/semver/v4"
)

const (
	// Unreleased tag is defined in keepachangelog 1.0.0.
	// https://keepachangelog.com/en/1.0.0/
	Unreleased        = "[Unreleased]"
	unreleasedHeading = "## [Unreleased]"
)

// Version transforms 0.1.0 to [0.1.0].
// Returns error when given input is not following SemVar.
func Version(in string) (string, error) {
	githubStyleSemver := false

	if len(in) > 0 && in[0] == 'v' {
		in = in[1:]
		githubStyleSemver = true
	}

	v, err := semver.Make(in)

	if err != nil {
		return "", errors.New("given version is not compatible with Semantic Versioning")
	}

	if githubStyleSemver == true {
		return fmt.Sprintf("[v%s]", v.String()), nil
	} else {
		return fmt.Sprintf("[%s]", v.String()), nil
	}
}

// To returns document replaced with given version.
func To(doc string, ver string) (string, error) {
	if doc == "" {
		return "", errors.New("given document is empty")
	}

	v, err := Version(ver)

	if err != nil {
		return "", err
	}

	// Check for diff and return the original when no changes.
	diff, err := Show(doc, Unreleased)

	if err != nil {
		return "", err
	}

	if len(diff) == 0 {
		return doc, nil
	}

	count := strings.Count(doc, unreleasedHeading)

	if count == 0 {
		return "", fmt.Errorf("given document does not contain %s tag", Unreleased)
	}

	if count > 1 {
		return "", fmt.Errorf("given document contains more than 1 %s tags", Unreleased)
	}

	template := strings.Join([]string{
		Unreleased,
		"",
		fmt.Sprintf("## %s - %s", v, time.Now().Format("2006-01-02")),
	}, "\n")

	return strings.Replace(doc, Unreleased, template, 1), nil
}

// Show returns changes of given version.
func Show(doc string, ver string) ([]string, error) {
	outs := []string{}

	if doc == "" {
		return outs, errors.New("given document is empty")
	}

	v := Unreleased

	if ver != Unreleased {
		ver, err := Version(ver)

		if err != nil {
			return outs, err
		}

		v = ver
	}

	found := false

	for _, line := range strings.Split(doc, "\n") {
		p := fmt.Sprintf("## %s", v)

		if !found && !strings.HasPrefix(line, p) {
			continue
		}

		// Mark as found and go to next cursor.
		if !found {
			found = true
			continue
		}

		// Finish when next heading found.
		if strings.HasPrefix(line, "## ") {
			break
		}

		outs = append(outs, line)
	}

	if len(outs) == 0 {
		return outs, nil
	}

	// Remove first line if empty.
	if outs[0] == "" {
		outs = append(outs[:0], outs[1:]...)
	}

	// Remove last line if empty.
	if len(outs) > 1 && outs[len(outs)-1] == "" {
		outs = outs[:len(outs)-1]
	}

	return outs, nil
}

// Latest returns the latest version stored in document.
// This operation simply matches to the first h2 header.
func Latest(doc string) (string, error) {

	re := regexp.MustCompile(`## \[([v]?\d*\.\d*\.\d*)\]`)

	for _, line := range strings.Split(doc, "\n") {
		got := re.FindStringSubmatch(line)

		if len(got) != 2 {
			continue
		}

		return got[1], nil
	}

	return "", errors.New("not found")
}

// LatestAny returns the latest [version] stored in document.
// Unlike `Latest` which follows Semantic Versioning, this function parse arbitary string
// excluding `## [Unreleased]` and string with blank spaces.
// This operation simply matches to the first h2 header.
func LatestAny(doc string) (string, error) {

	re := regexp.MustCompile(`## \[(.*)\]`)

	for _, line := range strings.Split(doc, "\n") {
		got := re.FindStringSubmatch(line)

		if len(got) != 2 {
			continue
		}

		if got[1] == "Unreleased" {
			continue
		}

		if strings.Contains(got[1], " ") {
			continue
		}

		return got[1], nil
	}

	return "", errors.New("not found")
}
