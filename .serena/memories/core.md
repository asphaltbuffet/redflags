# Core

`redflags` — a Go static analysis linter that enforces consistent long/short flag name pairings in CLI applications using cobra/pflag.

## Source Map

- `analyzer.go` — package `redflags`; exports `New(*Options) *analysis.Analyzer`. Contains `flagMappings`, `Options`, `run`, `visit`.
- `analyzer_test.go` — uses `analysistest.Run` against `testdata/` subdirs.
- `cmd/redflags/main.go` — standalone binary entry point; uses `singlechecker.Main`.
- `testdata/fn_valid/example.go` — valid flag usage (no diagnostics expected).
- `testdata/fn_invalid_short/example.go` — invalid short flags; `// want` annotations mark expected diagnostics.

## Key Invariants

- Package is `redflags` (library); binary is `cmd/redflags` (main).
- `flagMappings` is the canonical source of truth for long↔short pairings.
- Analyzer detects calls to methods ending in `P` (pflag convention); `VarP` variants shift arg index by 1.
- The `-V` flag is overridden in main.go to print version (avoids conflict with `version`→`V` mapping).

See `mem:tech_stack`, `mem:suggested_commands`, `mem:conventions`, `mem:task_completion`.
