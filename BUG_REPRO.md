# Bug Reproduction

Version parsing rejects valid `v`-prefixed input, accepts malformed ranges, and swallows errors through application and matcher layers; recheck also ignores cancellation. Run the eight `verify_cmds` in `collection.json`; all targeted range, propagation, and context tests fail. The affected symbols are in vulnerability domain, application, and infrastructure packages.
