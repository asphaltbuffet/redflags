# redflags — Core

Go static analysis linter enforcing consistent long/short cobra/pflag flag name pairings.
Implements `golang.org/x/tools/go/analysis` interface.

## Source map

- `analyzer.go` — library root; exports `New(*Options) *analysis.Analyzer`
- `analyzer_test.go` — tests via `analysistest.Run` + `// want` annotations in testdata
- `cmd/redflags/main.go` — standalone binary (`singlechecker.Main`)
- `testdata/<scenario>/example.go` — real Go source used as linter inputs
- `go.mod`, `go.sum`, `gomod2nix.toml` — module + Nix lockfile (kept in sync)

## Key invariants

- `flagMappings` is bidirectional: both `"verbose"→"v"` and `"v"→"verbose"` are stored
- `visit` checks method name suffix `P` / `VarP`; `VarP` shifts arg index by 1 (pointer receiver first)
- Diagnostics use `pass.Reportf`; no diagnostic = silently return
- `analysistest.Run` matches `// want <regex>` annotations to expected diagnostics

See `mem:tech_stack`, `mem:conventions`, `mem:suggested_commands`, `mem:task_completion`
