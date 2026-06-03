# Task Completion

When a coding task is done, run:

1. `mise run test` — must pass
2. `mise run lint` — must pass (golangci-lint --fix included)

For dependency changes also run `mise run mod-tidy`.
