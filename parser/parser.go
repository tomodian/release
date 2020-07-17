package parser

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/blang/semver/v4"
)

// List of tags.
const (
	Unreleased        = "[Unreleased]"
	unreleasedHeading = "## [Unreleased]"
)

// Version transforms 0.1.0 to [0.1.0].
// Returns error when given input is not following SemVar.
func Version(in string) (string, error) {
	v, err := semver.Make(in)

	if err != nil {
		return "", errors.New("given version is not compatible with Semantic Versioning")
	}

	return fmt.Sprintf("[%s]", v.String()), nil
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
