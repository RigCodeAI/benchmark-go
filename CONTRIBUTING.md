# Contributing

Contributions that improve case correctness, scanner interoperability,
documentation, framework coverage, and independent validation are welcome.

Before opening a pull request:

```bash
make verify
```

Case and truth changes must follow `docs/adding-a-test.md`. In particular, every
new category needs vulnerable, safe, unknown, and unsupported controls. Expected
results cannot be changed simply because a scanner disagrees with them.

Pull requests should state:

- the security claim being added or corrected;
- the affected CWE and exact dependency/runtime coordinate;
- why the safe case is authoritative;
- whether public accuracy or qualification behavior changes; and
- how the change was independently validated.

By contributing, you agree that your contribution is licensed under this
repository's MIT license.
