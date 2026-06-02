# Tech Stack

- **Language**: Go 1.22
- **Analysis framework**: `golang.org/x/tools/go/analysis` — `analysis.Analyzer`, `inspect.Analyzer`, `analysistest`
- **CLI framework detected in linted code**: `github.com/spf13/cobra`, `github.com/spf13/pflag` (these are in testdata, not go.mod)
- **Task runner**: `mise` (mise.toml defines all tasks)
- **Linter**: `golangci-lint` v2.8.x (via aqua)
- **Test runner**: `gotestsum` (via aqua)
- **Mock generator**: `mockery` v3 (via aqua)
- **Release**: `goreleaser`
- **Changelog**: `changie`
- **Nix lockfile**: `gomod2nix` — `gomod2nix.toml` must be kept in sync with `go.sum` via `mise run nix-hash`
