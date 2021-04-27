# Release

![Test on Linux](https://github.com/tomodian/release/workflows/Test%20on%20Linux/badge.svg?branch=develop)
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

This command will show you all CHANGELOG.md files recursively.

    release target
    release target --dir path/to/entrypoint
    release t -d path/to/entrypoint

### See unreleased changes

Grab `[Unreleased]` sections of all CHANGELOG.md files recursively.

    release unreleased
    release unreleased --dir path/to/entrypoint
    release u -d path/to/entrypoint

### See previous versions

    release show -v 0.1.0
    release show -v 0.1.0 --dir path/to/entrypoint
    release s -v 0.1.0 -d path/to/entrypoint

### Show the latest released version in current directory

    release latest
    release l

### Bump all [Unreleased] sections to given version

By default, `release` will ask you for confirmation.

    release to -v 0.2.0

    # Targets
    ## .github/workflows/CHANGELOG.md
    ## CHANGELOG.md
    ✔ Enter `yes` to update all CHANGELOGs to version [0.8.0]: yes

If you want to integrate with CI pipeline, use `--force` or `-f`.

    release to -v 0.2.0 --force

    # Targets
    ## .github/workflows/CHANGELOG.md --> ✅
    ## CHANGELOG.md --> ✅
    Done👍

### See next release version

`release` will suggest you the next available version.

    release next

    Latest released version: 0.8.0

    Suggestions for next release:
       - Major / Release --> 1.0.0
       - Minor / Feature --> 0.9.0
       - Patch / Hotfix  --> 0.8.1

For CI integrations, add `--type` flag.
The words `major`, `minor` and `patch` comes from [Semantic Versioning 2.0.0](https://semver.org) idiom.

    release next --type major
    1.0.0

    release next --type minor
    0.9.0

    release next --type patch
    0.8.1

[GitFlow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) idiom is also supported.

    release next --type release
    1.0.0

    release next --type feature
    0.9.0

    release next --type hotfix
    0.8.1

## Development

### Run

    make run

### Test

    make test

### Build

    make build

## Issues?

### MacOS "developer cannot be verified" error

Note MacOS Catalina will warn you when executing this binary in command-line.
Follow these steps to give your permission.

1. In the Finder on your Mac, unzip the `release` app and open the `release` binary
2. MacOS will prompt you to enable the binary, so answer yes
3. Now you can use the binary from command-line.

## License

[Mozilla Public License v2.0](LICENSE)
