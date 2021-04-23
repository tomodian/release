package parser

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
