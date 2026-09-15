# Release

[![Go Reference](https://pkg.go.dev/badge/github.com/tomodian/release.svg)](https://pkg.go.dev/github.com/tomodian/release)
[![Test on Linux](https://github.com/tomodian/release/actions/workflows/test.linux.yml/badge.svg)](https://github.com/tomodian/release/actions/workflows/test.linux.yml)
[![Release](https://github.com/tomodian/release/actions/workflows/release.yml/badge.svg)](https://github.com/tomodian/release/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/tomodian/release/branch/develop/graph/badge.svg)](https://codecov.io/gh/tomodian/release)

[![GitHub All Releases](https://img.shields.io/github/downloads/tomodian/release/total?style=social)](https://github.com/tomodian/release/releases)

A small command-line utility to manage CHANGELOG.md written in [keepachangelog.com](https://keepachangelog.com) format.

Works nicely on any sized Git repository, even awesome on [Monorepo](https://en.wikipedia.org/wiki/Monorepo).

## Installation

Please download ZIP archive from [releases](https://github.com/tomodian/release/releases) page.

## How it works

Run `release` to show full list of commands and flags.

### List all CHANGELOG.md

`release target` will show you all CHANGELOG.md files recursively.

```bash
release target
release target --dir path/to/entrypoint
release t -d path/to/entrypoint
```

### See unreleased changes

`release unreleased` will grab `[Unreleased]` sections of all CHANGELOG.md files recursively.

```bash
release unreleased
release unreleased --dir path/to/entrypoint
release u -d path/to/entrypoint
```

### See previous versions

`release show` will output all previous version histories.

```bash
release show -v 0.1.0
release show -v 0.1.0 --dir path/to/entrypoint
release s -v 0.1.0 -d path/to/entrypoint
```

### Show the latest released version in current directory

```bash
release latest
release latest --newline=false
release l
```

### Bump all [Unreleased] sections to given version

By default, `release to -v X.Y.Z` will ask you for confirmation.

```bash
release to -v 0.2.0

# Targets
## .github/workflows/CHANGELOG.md
## CHANGELOG.md
✔ Enter `yes` to update all CHANGELOGs to version [0.8.0]: yes
```

If you want to integrate with CI pipeline, use `--force` or `-f`.

```bash
release to -v 0.2.0 --force

# Targets
## .github/workflows/CHANGELOG.md --> ✅
## CHANGELOG.md --> ✅
Done👍
```

### See next release version

`release next` will suggest you the next available version.

```bash
release next

Latest released version: 0.8.0

Suggestions for next release:
- Major / Release --> 1.0.0
- Minor / Feature --> 0.9.0
- Patch / Hotfix  --> 0.8.1
```

For CI integrations, add `--type` flag.
The words `major`, `minor` and `patch` comes from [Semantic Versioning 2.0.0](https://semver.org) idiom.

```bash
release next --type major
1.0.0

release next --type minor
0.9.0

release next --type patch
0.8.1
```

Note this command will not add newline when `--type` flag is specified.
Use `--newline` flag if you prefer to see the newline.

```bash
release next --type major --newline
```

[GitFlow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) idiom is also supported.

```bash
release next --type release
1.0.0

release next --type feature
0.9.0

release next --type hotfix
0.8.1
```

### Github-style semver `vx.y.z`

The tool also supports [Github-style semver](https://semver.org/#is-v123-a-semantic-version):

```bash
release show -v v0.1.0
release to -v v0.2.0
```

## Development

### Run

```bash
make run
```

### Test

```bash
make test
```

### Build

```bash
make build
```

## License

[Mozilla Public License v2.0](LICENSE)
