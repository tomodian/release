# Release

[![GitHub All Releases](https://img.shields.io/github/downloads/tomodian/release/total?style=social)](https://github.com/tomodian/release/releases)
![Test on Linux](https://github.com/tomodian/release/workflows/Test%20on%20Linux/badge.svg?branch=develop)
[![codecov](https://codecov.io/gh/tomodian/release/branch/develop/graph/badge.svg)](https://codecov.io/gh/tomodian/release)

A small command-line utility to manage CHANGELOG.md written in [keepachangelog.com](https://keepachangelog.com) format.

Works nicely on [Monorepo](https://en.wikipedia.org/wiki/Monorepo).

## Installation

Please download ZIP archive from [releases](https://github.com/tomodian/release/releases) page.

### MacOS "developer cannot be verified" error

Note MacOS Catalina will warn you when executing this binary in command-line.
Follow these steps to give your permission.

1. In the Finder on your Mac, unzip the `release` app and open the `release` binary
2. MacOS will prompt you to enable the binary, so answer yes
3. Now you can use the binary from command-line.

## How it works

Run `release` to show full list of commands and flags.

### List all CHANGELOG.md in directory

    release target

### See unreleased changes

    release next

### See previous versions

    release show -v 0.1.0

## Development

### Run

    make run

### Test

    make test

### Build

    make build

## License

[Mozilla Public License v2.0](LICENSE)
