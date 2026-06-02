# redflags
a linter for flags in go cli applications

[![Release](https://img.shields.io/github/v/release/asphaltbuffet/redflags?style=flat-square)](https://github.com/asphaltbuffet/redflags/releases)
[![go.mod](https://img.shields.io/github/go-mod/go-version/asphaltbuffet/redflags)](go.mod)
[![License](https://img.shields.io/github/license/asphaltbuffet/redflags?style=flat-square)](LICENSE)
![Changelog](https://img.shields.io/badge/Common%20Changelog-blue?style=flat-square&link=CHANGELOG.md)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/asphaltbuffet/redflags/codeql.yml?style=flat-square)
[![CodeQL](https://github.com/asphaltbuffet/redflags/workflows/CodeQL/badge.svg?style=flat-square)](https://app.codecov.io/gh/asphaltbuffet/redflags)
[![wakatime](https://wakatime.com/badge/github/asphaltbuffet/redflags.svg?style=flat-square)](https://wakatime.com/badge/github/asphaltbuffet/redflags)

## Usage

### As a golangci-lint plugin

golangci-lint module plugins require a custom build of the linter. Create a `.custom-gcl.yml` in your project:

```yaml
version: v2.9.0  # match your installed golangci-lint version
plugins:
  - module: github.com/asphaltbuffet/redflags
    version: v0.1.0  # pin to a release
```

Then build your custom `golangci-lint` binary:

```sh
golangci-lint custom
```

Enable and configure `redflags` in your `.golangci.yml`:

```yaml
version: "2"

linters:
  enable:
    - redflags
  settings:
    custom:
      redflags:
        type: module
        description: Enforces consistent long/short flag name pairings in cobra/pflag CLIs.
        settings:
          use-defaults: true      # include built-in pairs (help↔h, verbose↔v, output↔o)
          mappings:
            - long: deploy
              short: d
            - long: config
              short: c
```

The built-in default pairs are `help↔h`, `verbose↔v`, and `output↔o`. Set `use-defaults: false` to use only your own `mappings`.

### As a standalone binary

Download a release binary from the [releases page](https://github.com/asphaltbuffet/redflags/releases) and run it directly against a package:

```sh
redflags ./...
```

## Community

### Code of Conduct

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## References
- https://golangci-lint.run/contributing/new-linters/
- https://disaev.me/p/writing-useful-go-analysis-linter/
- http://goast.yuroyoro.net/
