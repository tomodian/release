package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/facette/natsort"
)

// SemanticVersion represents major/minor/patch numbers.
type SemanticVersion struct {
	Major int
	Minor int
	Patch int
}

// NewSemanticVersion takes arbitary string, parse and return struct.
func NewSemanticVersion(given string) (*SemanticVersion, error) {

	got := strings.Split(given, ".")

	if len(got) != 3 {
		return nil, errors.New("input must have 3 integers concatenated by period (e.g. 1.2.3")
	}

	out := SemanticVersion{}

	{
		var err error

		if out.Major, err = CastVersion("major", got[0]); err != nil {
			return nil, err
		}

		if out.Minor, err = CastVersion("minor", got[1]); err != nil {
			return nil, err
		}

		if out.Patch, err = CastVersion("patch", got[2]); err != nil {
			return nil, err
		}
	}

	return &out, nil
}

// IsEqual compares internal version with given version.
func (c SemanticVersion) IsEqual(in *SemanticVersion) bool {
	if in == nil {
		return false
	}

	return c.Major == in.Major && c.Minor == in.Minor && c.Patch == in.Patch
}

// IsGreater compares internal version with given version.
func (c SemanticVersion) IsGreater(in *SemanticVersion) bool {
	if in == nil {
		return false
	}

	switch {
	case c.Major < in.Major:
		return true

	case c.Major == in.Major && c.Minor < in.Minor:
		return true

	case c.Major == in.Major && c.Minor == in.Minor && c.Patch < in.Patch:
		return true
	}

	return false
}

// String decorator.
func (c SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", c.Major, c.Minor, c.Patch)
}

// Increment version according to given type.
// The type must be one of Major/Minor/Patch.
//
// Given 1.2.3 as example:
// - Increment(Major) --> 2.2.3
// - Increment(Minor) --> 1.3.3
// - Increment(Patch) --> 1.2.4
func (c SemanticVersion) Increment(in VersionType) SemanticVersion {
	switch in {
	case MajorVersion:
		return SemanticVersion{Major: c.Major + 1, Minor: 0, Patch: 0}

	case MinorVersion:
		return SemanticVersion{Major: c.Major, Minor: c.Minor + 1, Patch: 0}
	}

	return SemanticVersion{Major: c.Major, Minor: c.Minor, Patch: c.Patch + 1}
}

// CastVersion parse and set a string into given struct.
func CastVersion(name, val string) (int, error) {
	const failcode = -1

	if name == "major" {
		if val[0] == 'v' {
			val = val[1:]
		}
	}

	i, err := strconv.Atoi(val)

	if err != nil {
		return failcode, fmt.Errorf("%s segment must be integer", name)
	}

	if i < 0 {
		return failcode, fmt.Errorf("%s segment must be greater than zero", name)
	}

	return i, nil
}

// SortVersions in ascending order with dumb algorithm.
func SortVersions(in []SemanticVersion) []SemanticVersion {

	vers := []string{}

	for _, v := range in {
		vers = append(vers, v.String())
	}

	natsort.Sort(vers)

	out := []SemanticVersion{}

	for _, v := range vers {
		got, _ := NewSemanticVersion(v)
		out = append(out, *got)
	}

	return out
}
