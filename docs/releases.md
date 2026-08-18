# Releases and provenance

Release tags use `vMAJOR.MINOR.PATCH`. The release workflow verifies the benchmark,
builds the independent scorer, archives the source at the tagged commit, publishes
`SHA256SUMS`, and requests a GitHub artifact attestation for the artifacts.

Consumers should pin both the tag and commit digest. A release tag alone is not a
sufficient product qualification identity. Downstream Rig qualification records
the immutable benchmark commit and verifies benchmark-owned scoring independently.

The signing/attestation workflow is release infrastructure, not evidence that any
scanner achieved a particular score.
