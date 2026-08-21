# Bug Reproduction

Wrapped sentinel errors are classified by message text instead of the error chain, so HTTP mapping falls back to 422. Run the four `verify_cmds` in `collection.json`; wrapped not-found, conflict, duplicate, and unauthorized cases fail their expected status assertions. The affected symbols are in `internal/platform/not_found.go`, `conflict.go`, `duplicate.go`, and `unauthorized.go`.
