# Conventions

- Tests use `analysistest.Run` — NOT standard `t.Error`. Expected diagnostics are `// want <regex>` comments in testdata files.
- Each testdata scenario is its own package under `testdata/<scenario>/example.go`.
- Adding a flag pair: add both directions to `flagMappings`, add testdata with `// want` annotation.
- `//nolint:errcheck` used on type assertions the framework guarantees (e.g. inspector results).
- `//nolint:gochecknoglobals` on package-level immutable vars.
- No diagnostics for unknown/unmapped flags — only report when a known pairing is violated.
