# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.1.0-rc2] - 2023-02-28

### Features

- Adding `--raw` flag for export function which will output scale functions as `wasm` files directly instead of encoded `.scale` files
- Adding aliases for `scale push`, `scale run`, `scale new`, and `scale build` commands

### Fixes

- Returning clearer error messages when the names or tags of functions are invalid

### Dependencies



## [v0.1.0-rc1] - 2023-02-20

### Features

- Initial release of the Scale CLI
- Added support for functions written in `Go` and `Rust`
- Added support for running `http` functions locally
- Added support for pushing and pulling functions from the scale registry

[unreleased]: https://github.com/loopholelabs/scale-cli/compare/v0.1.0-rc2...HEAD
[v0.1.0-rc2]: https://github.com/loopholelabs/scale-cli/compare/v0.1.0-rc2
[v0.1.0-rc1]: https://github.com/loopholelabs/scale-cli/compare/v0.1.0-rc1
