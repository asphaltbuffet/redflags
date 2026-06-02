# Conventions

- New flag mappings go in `flagMappings` in `analyzer.go` — bidirectional (both `long→short` and `short→long`).
- Test cases live in `testdata/<scenario>/example.go`; use `// want <backtick-regex>` annotations for expected diagnostics.
- `analysistest.TestData()` returns the directory containing `testdata/`; subdirs are passed as patterns.
- `VarP` methods (e.g. `BoolVarP`) take a pointer as first arg, shifting long/short name args to index 1/2.
- `Options.ShortToLong` / `LongToShort` fields exist but are not yet wired into the `visit` logic (as of initial implementation).
- **Go version must stay at `go 1.25.6`** — upgrading beyond 1.25.x requires significant golangci-lint compatibility work. Do not bump without explicitly resolving that first.
