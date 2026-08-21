# Bug Reproduction

Webhook outbox enqueue and drain operations retain nested payload/header aliases and queue backing storage. Run the five `verify_cmds` in `collection.json`; caller mutations change queued or drained events and storage is retained after drain. The affected symbols are in `internal/webhook/application`.
