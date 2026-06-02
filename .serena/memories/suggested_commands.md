# Suggested Commands

All tasks run via `mise run <task>` (or `mise run` to list). Do NOT use `git` — use `jj` for VCS.

| Task | Command |
|---|---|
| Build binary | `mise run build` |
| Run tests | `mise run test` |
| Run single test | `go test -run TestFuncLinting/valid ./...` |
| Lint | `mise run lint` |
| Generate | `mise run generate` (or `mise run gen`) |
| Coverage HTML | `mise run cover` |
| Full dev pipeline | `mise run dev` |
| Mod tidy + nix hash | `mise run mod-tidy` |

Tests use `gotestsum` with `-race -covermode=atomic`; output goes to `bin/coverage.out`.
