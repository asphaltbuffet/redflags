# Suggested Commands

All commands require the Nix devshell. Prefix with `nix develop --command`:

| Task | Command |
|---|---|
| Run tests (force, bypassing mise cache) | `nix develop --command mise run --force test` |
| Lint | `nix develop --command mise run --force lint` |
| Build | `nix develop --command mise run --force build` |
| Full dev pipeline | `nix develop --command mise run --force dev` |
| Single test by name | `nix develop --command go test -run TestName ./...` |

Note: plain `go`, `mise`, `gotestsum` etc. are NOT on PATH outside `nix develop`.
The `mise run test` task is source-fingerprinted; use `--force` to bypass the cache.
