# Bug Reproduction

Storage adapter and application entry points ignore canceled contexts and can reach a nil repository. Run the five `verify_cmds` in `collection.json`; canceled operations return nil or panic instead of propagating cancellation. The affected symbols are in `internal/storage/adapter` and `internal/storage/application`.
