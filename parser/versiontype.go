package parser

import "errors"

// VersionType corresponds to Major.Minor.Patch in Semantic Versioning 2.0.0.
// https://semver.org
type VersionType string

// String decorator.
func (c VersionType) String() string {
	return string(c)
}

// Exposed keys of version types.
const (
	MajorVersion VersionType = "major"
	MinorVersion VersionType = "minor"
	PatchVersion VersionType = "patch"
)

// AliasedVersion returns the original version type or error.
// GitFlow idiom is currently available.
func AliasedVersion(in string) (VersionType, error) {

	switch in {

	case MajorVersion.String(), "release":
		return MajorVersion, nil

	case MinorVersion.String(), "feature":
		return MinorVersion, nil

	case PatchVersion.String(), "hotfix":
		return PatchVersion, nil
	}

	return "", errors.New("given alias is not in list")
}
