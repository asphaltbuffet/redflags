# Task Completion

Run before considering a coding task done:

1. `mise run lint` — golangci-lint with --fix
2. `mise run test` — full test suite with race detector
3. If go.mod/go.sum changed: `mise run mod-tidy` (also updates gomod2nix.toml)
