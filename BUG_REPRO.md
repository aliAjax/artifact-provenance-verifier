# Bug Reproduction

`LeaseJob` returns before its callback completes and drops callback errors; the lease adapter ignores cancellation and owner checks. Run the five `verify_cmds` in `collection.json`; each test fails in the worker lease lifecycle. The affected symbols are in `internal/worker/application/lease.go` and `internal/worker/adapter/lease.go`.
