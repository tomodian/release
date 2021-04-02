# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.5.1] - 2021-04-02

### Fixed

- Excluded darwin/386 arch due to deprecation

## [0.5.0] - 2021-03-25

### Changed

- Render Markdown as default style

## [0.4.0] - 2020-10-07

### Added

- More directories to exclude from grep
- Latest task to show the latest released version

## [0.3.2] - 2020-07-20

### Added

- Some more tests

## [0.3.1] - 2020-07-20

### Added

- Makefile task to see coverage report
- Some tests

## [0.3.0] - 2020-07-17

### Added

- Codecov reports and badges

### Changed

- Show CHANGELOG.md in relative path

## [0.2.2] - 2020-07-17

### Added

- Ignored common vendor directories

### Changed

- Moved utils.Glob to files

## [0.2.1] - 2020-07-17

### Added

- Command examples on README

## [0.2.0] - 2020-07-17

### Added

- File writer to persist changes
- More tests for utils

### Changed

- Moved utils.ReadFile to files.Read

## [0.1.7] - 2020-07-17

### Changed

- Switched to use 'hub release edit' command to attach files

## [0.1.6] - 2020-07-17

### Fixed

- Typo on hub release

## [0.1.5] - 2020-07-17

### Fixed

- Explicitly added Go binary path

## [0.1.4] - 2020-07-17

### Fixed

- Trying to download gox by directly calling go get

## [0.1.3] - 2020-07-17

### Fixed

- Setup actions/setup-go@v2 before calling Build step

## [0.1.2] - 2020-07-17

### Fixed

- Use actions/setup-go@v2 for build

## [0.1.1] - 2020-07-17

### Fixed

- Added installation step on release

## [0.1.0] - 2020-07-17

### Added

- GitHub Actions release.yml
- GitHub Actions test.linux.yml
- License
- Executable bundler for \*nix and Windows
- Parser and utils
