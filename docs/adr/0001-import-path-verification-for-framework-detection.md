# Import-path verification for framework detection

The analyzer detects flag-registration calls by checking method name suffixes (e.g. `BoolP`, `StringVarP`). This heuristic alone would fire on any package that happens to define methods with those names, producing false positives. We resolved this by verifying the receiver's import path via `pass.TypesInfo` — confirming it resolves to `github.com/spf13/cobra` or `github.com/spf13/pflag` before reporting a diagnostic.

The alternative (suffix matching only) is simpler but erodes trust immediately when it misfires. The import-path approach also establishes the extension pattern for future frameworks: adding kong, cli, etc. means adding import paths to a verified set, not redesigning detection.
