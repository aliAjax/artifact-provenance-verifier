# Bug Reproduction

Unit-of-work and SQL store paths mask operation errors, repeat callbacks, lose contexts, and dereference nil receivers. Run the seven `verify_cmds` in `collection.json`; the transaction and ping assertions fail or panic. The affected symbols are in `internal/storage/domain/transaction.go` and `internal/storage/adapter/postgres.go`.
