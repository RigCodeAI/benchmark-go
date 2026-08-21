# Scanner contracts

BenchmarkGo has two deliberately separate contracts.

## Public accuracy contract

Any scanner may submit SARIF 2.1.0, CSV, or JSON conforming to
`schemas/scanner-results-v1.schema.json`. Findings identify a standard CWE or
Go-specific security category and at least one public case identity: route,
case-bearing repository location, or stable case ID. No Sivere
publication, transcript, runtime coordinate, or evidence grade is required.

```bash
./scoreBenchmark.sh --results results/my-tool.sarif --output-dir results/my-tool
```

The scorer matches submitted findings to the public case catalog. Absence on a
vulnerable case is an FN. A finding on a safe case is an FP. Unmapped findings
are also FP. Duplicate reports do not improve the score.

## High-assurance qualification contract

Products claiming runtime evidence and complete coverage may additionally emit
`schemas/qualification-evidence-v1.schema.json`. This contract binds observations
to repository, runtime, framework, publication, coverage, transcripts, and
authenticated readback. The benchmark-owned scorer alone compares it with truth.

```bash
cargo run --release --locked -- score \
  --truth truth-v2.json \
  --evidence /immutable/scanner-evidence.json \
  --output /immutable/qualification.json
```

`--held-out` may be repeated. `--require-promotion` exits 30 until the exact score,
closed envelope, coordinate matrix, and independent held-out gates pass.

The scanner must never read benchmark truth during analysis. Product adapters may
read the public catalog, but truth comparison and score calculation remain owned
by this repository.
