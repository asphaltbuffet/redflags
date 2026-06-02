# Conventions

- New flag mappings go in `defaultFlagPairs` in `analyzer.go` — bidirectional (both `long→short` and `short→long`).
- Test cases live in `testdata/<scenario>/example.go`; use `// want <backtick-regex>` annotations for expected diagnostics.
- `analysistest.TestData()` returns the directory containing `testdata/`; subdirs are passed as patterns.
- `VarP` methods (e.g. `BoolVarP`) take a pointer as first arg, shifting long/short name args to index 1/2.
- Built-in defaults are `help↔h`, `verbose↔v`, `output↔o` (no `version↔V`).
- **Go version must stay at `go 1.25.6`** — upgrading beyond 1.25.x requires significant golangci-lint compatibility work. Do not bump without explicitly resolving that first.
- **Always use `fd -H`** — plain `fd` silently skips hidden files, causing false "file not found" results. `-H` includes hidden files while still respecting `.gitignore`.
- The `exhaustruct` linter requires all struct fields to be listed explicitly in struct literals.
