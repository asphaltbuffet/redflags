# Suggested Commands

All tasks via `mise run <task>`:

| Task | Purpose |
|---|---|
| `mise run test` | Run test suite (gotestsum, race detector, coverage) |
| `mise run lint` | golangci-lint with fix |
| `mise run build` | Build binary to dist/ |
| `mise run dev` | Full pipeline: generate → lint → test → snapshot |
| `mise run cover` | Generate HTML coverage report |
| `mise run mod-tidy` | go mod tidy + gomod2nix generate |

Run single test: `go test -run TestFuncLinting/valid ./...`

VCS: use `jj` (not `git`) for all version control operations.
