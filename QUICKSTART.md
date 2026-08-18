# BenchmarkGo quickstart

## 1. Verify prerequisites

```bash
go version
cargo --version
```

The launcher requests Go 1.25.12 for `net/http` and Go 1.26.5 for Gin using
Go's standard `GOTOOLCHAIN` support.

## 2. Start a benchmark application

```bash
./runBenchmark.sh net-http
```

Leave it running at `http://127.0.0.1:8080`. Use
`./runBenchmark.sh gin` to exercise the Gin application instead.

## 3. Run a scanner

- SAST: scan `apps/net-http-product` or `apps/gin-product`.
- DAST/IAST: scan the running application.
- Hybrid tools may do both.

Export findings as SARIF 2.1.0, the JSON schema under `schemas`, or the CSV
columns documented in `docs/scanner-integration.md`.

## 4. Create a scorecard

```bash
./scoreBenchmark.sh \
  --results /path/to/scanner-results.sarif \
  --output-dir results/my-tool-1.2.3
```

Open `results/my-tool-1.2.3/scorecard.html` in a browser. The JSON and CSV files
beside it are intended for automation and comparison.

## 5. Verify the benchmark itself

```bash
./verifyBenchmark.sh
```

This verifies the independent scorer, generated case catalog, all example input
formats, the Go language corpus, and both public applications.
