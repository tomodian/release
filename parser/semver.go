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

	if c.Major == in.Major && c.Minor == in.Minor && c.Patch == in.Patch {
		return true
	}

	return false
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

// CastVersion parse and set a string into given struct.
func CastVersion(name, val string) (int, error) {
	const failcode = -1

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
