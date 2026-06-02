# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Is

`redflags` is a Go static analysis linter that enforces consistent long/short flag name pairings in cobra/pflag CLI applications. It implements the `golang.org/x/tools/go/analysis` interface, making it usable standalone or as a golangci-lint plugin.

## Commands

All tasks are defined in `mise.toml` and run via `mise run <task>`.

| Task | Command |
|---|---|
| Build binary | `mise run build` |
| Run tests | `mise run test` |
| Lint | `mise run lint` |
| Generate | `mise run generate` |
| Full dev pipeline | `mise run dev` |
| Coverage HTML | `mise run cover` |
| Mod tidy + update nix lockfile | `mise run mod-tidy` |

Run a single test by name:
```
go test -run TestFuncLinting/valid ./...
```

## Architecture

### Package layout

- `analyzer.go` (package `redflags`) — the library. Exports `New(*Options) *analysis.Analyzer`.
- `cmd/redflags/main.go` — standalone binary, wraps the analyzer with `singlechecker.Main`.
- `testdata/<scenario>/example.go` — real Go source files used as linter inputs in tests.

### How the analyzer works

`flagMappings` in `analyzer.go` is the bidirectional source of truth for valid long↔short pairings (e.g. `"verbose"→"v"` and `"v"→"verbose"`). The `visit` function inspects every `ast.CallExpr` for method calls ending in `P` (the pflag convention for calls that accept a short flag). For `VarP` variants, argument indices shift by 1 because the first arg is a pointer receiver.

### Adding a new flag pair

Add both directions to `flagMappings`, then add a test case in `testdata/` with a `// want` annotation.

### Testing pattern

Tests use `analysistest.Run`, which drives the analyzer against real Go source in `testdata/` and matches `// want <regex>` comment annotations to expected diagnostics — not standard `t.Error` assertions.

## Nix

`gomod2nix.toml` is a Nix dependency lockfile derived from `go.sum`. It must stay in sync. `mise run mod-tidy` handles this automatically (calls `gomod2nix generate` as a post-step).

## VCS

Use `jj` (jujutsu), not `git`.
