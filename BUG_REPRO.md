# Bug Reproduction

Typed-nil senders, nil HTTP clients or receivers, and a nil circuit receiver are dereferenced before validation. Run the seven `verify_cmds` in `collection.json`; the tests report configuration failures or nil-pointer panics. The affected symbols are in the webhook application, adapter, and infrastructure packages.
