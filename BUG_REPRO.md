# Bug Reproduction

Verification cache and clone helpers return shallow copies, allowing `Checks` slices and `Details` maps to alias caller and cache state. Run the six `verify_cmds` in `collection.json`; mutations after Put or Get change subsequent results. The affected symbols are in `internal/platform/model.go`, `internal/verification/infrastructure/cache.go`, and `internal/verification/application/clone.go`.
