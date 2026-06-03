# Context

## Glossary

**Flag Pair**
A declared relationship between a long flag name and a short flag name. e.g. `--verbose` / `-v`. A flag pair is 1:1 — one long name maps to exactly one short name.

**Default Mappings**
The built-in set of flag pairs enforced when `use-defaults` is `true` (the default). Intentionally small — only pairs that are nearly universal across CLI tools. Current set: `help ↔ h`, `verbose ↔ v`, `output ↔ o`.

**Custom Mappings**
User-defined flag pairs declared in `.golangci.yml` under the linter's settings block. Always additive relative to the default mappings. Setting `use-defaults: false` replaces default mappings entirely with the custom set.

**Pairing Correctness**
The check redflags enforces: if a long flag name appears in the active mapping set, its short flag must match the mapped value — and vice versa. Both directions are always enforced. Pairing correctness does not check whether a flag *has* a short form at all (that is a Presence Requirement, deferred to a future version).

**Presence Requirement**
A future check (not in v1): if a long flag that has a known mapping exists in a command, its paired short flag must also be declared. Requires scanning all flags in a command as a unit rather than one call at a time.

**Supported Framework**
A CLI library whose flag-registration calls redflags understands. v1 supports `github.com/spf13/cobra` and `github.com/spf13/pflag`. Framework detection is verified via import-path resolution, not method-name heuristics alone.

**`P`-variant call**
A flag-registration method call whose name ends in `P` (e.g. `BoolP`, `StringVarP`) — the pflag/cobra convention for calls that accept both a long name and a short name. `VarP` variants shift the long/short argument indices by one due to the leading pointer argument.
