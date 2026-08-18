# Security policy

This repository intentionally contains vulnerable applications. Findings that are
part of the published case catalog are not security incidents.

Please privately report vulnerabilities that could affect benchmark users outside
the intended disposable applications—for example scorer input escapes, CI secret
exposure, unsafe release artifacts, or code execution during non-executing
inspection. Use GitHub's private vulnerability reporting or a private security
advisory for the repository rather than a public issue.

Never deploy the benchmark applications to a public or production environment.
Run them only on loopback or in an isolated container/network.
