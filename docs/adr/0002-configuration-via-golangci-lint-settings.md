# Configuration via golangci-lint settings block

User-defined flag pairs are configured via the linter's `settings:` block in `.golangci.yml` rather than a standalone config file. This keeps redflags consistent with every other golangci-lint plugin and avoids building a separate config-file reader.

The `Options` struct exposes two fields: `UseDefaults bool` (default `true`) and `Mappings []FlagPair`. Custom mappings are always additive over the defaults; setting `UseDefaults: false` replaces defaults entirely. A dedicated `ignore` list was considered but rejected — `use-defaults: false` plus an explicit `mappings` list covers the same need with fewer concepts.
