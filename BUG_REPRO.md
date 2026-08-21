# Bug Reproduction

The artifact status transition rules are inverted or bypassed across platform, domain, policy, and lifecycle layers. Run the eleven `verify_cmds` in `collection.json`; illegal restores are accepted, legal withdrawn restores are rejected, and invalid sources or targets pass. The affected symbols are named in the collection root-cause field.
