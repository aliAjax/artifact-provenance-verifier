# Bug Reproduction

`Memory` and `SnapshotStore` leak slices, maps, and byte-slice backing storage across calls. Run the eight `verify_cmds` in `collection.json`; each test fails with caller mutations visible in later reads. The affected symbols are in `internal/storage/infrastructure/memory.go` and `snapshot.go`.
