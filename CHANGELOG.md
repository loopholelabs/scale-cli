# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.1.4] - 2023-03-23

### Fixes

- Fixing a bug where new function analytics events would not be recorded
- Fixing bug where `new update available` text would not be displayed correctly

## [v0.1.3] - 2023-03-19

### Fixes

- Fixing a bug where Analytics events would not be properly pushed
- Fixing a bug where update analytics events would not be recorded

## [v0.1.2] - 2023-03-19

### Features

- Adding update check and auto update functionality to the CLI
- Adding --no-telemetry flag to disable telemetry collection (using PostHog)

### Dependencies

- Bumping `scale` version to `v0.3.15`
- Bumping `scale-http-adapters` version to `v0.3.8`
- Bumping `scale-signature-http` version to `v0.3.8`

## [v0.1.1] - 2023-03-12

### Fixes

- Fixing bugs where panics would occur if the user was not logged in 

## [v0.1.0] - 2023-03-10

### Features

- Various bug fixes and improvements for release

## [v0.1.0-rc2] - 2023-02-28

### Features

- Adding `--raw` flag for export function which will output scale functions as `wasm` files directly instead of encoded `.scale` files
- Adding aliases for `scale push`, `scale run`, `scale new`, and `scale build` commands

### Fixes

- Returning clearer error messages when the names or tags of functions are invalid

### Dependencies

- Bumping `golang.org/x/net` from `v0.5.0` to `v0.7.0`
- Bumping `golang.org/x/sys` from `v0.4.0` to `v0.5.0`
- Bumping `golang.org/x/term` from `v0.4.0` to `v0.5.0`
- Bumping `golang.org/x/text` from `v0.6.0` to `v0.7.0`
- Bumping `scale` from `v0.3.11` to `v0.3.12`
- Bumping `scale-http-adapters` from `v0.3.5` to `v0.3.6`

## [v0.1.0-rc1] - 2023-02-20

### Features

- Initial release of the Scale CLI
- Added support for functions written in `Go` and `Rust`
- Added support for running `http` functions locally
- Added support for pushing and pulling functions from the scale registry

[unreleased]: https://github.com/loopholelabs/scale-cli/compare/v0.1.4...HEAD
[v0.1.4]: https://github.com/loopholelabs/scale-cli/compare/v0.1.4
[v0.1.3]: https://github.com/loopholelabs/scale-cli/compare/v0.1.3
[v0.1.2]: https://github.com/loopholelabs/scale-cli/compare/v0.1.2
[v0.1.1]: https://github.com/loopholelabs/scale-cli/compare/v0.1.1
[v0.1.0]: https://github.com/loopholelabs/scale-cli/compare/v0.1.0
[v0.1.0-rc2]: https://github.com/loopholelabs/scale-cli/compare/v0.1.0-rc2
[v0.1.0-rc1]: https://github.com/loopholelabs/scale-cli/compare/v0.1.0-rc1
