# BenchmarkGo v1

BenchmarkGo is the closed qualification denominator for Rig's Go target. It has
39 categories and four controls per category: vulnerable, safe, unknown, and
unsupported. Its native Go corpus, ordinary product cases, controller cases,
source contracts, hostile repositories, resource budgets, schemas, truth, and
promotion rules are SHA-256 locked.

## What is implemented

- 156/156 executable language-corpus controls;
- 87 source-to-sink product cases: 29 vulnerable, 29 safe, and 29 unknown;
- 30 controller cases across authorization, authentication, CSRF, tenants,
  workflows, concurrency, and multi-service protocols;
- 11 source contracts and 27 foundation/hostile/coordinate cases;
- ordinary `rig run` execution with authenticated collection, sealed transcripts,
  native reduction/publication, and signed readback;
- native scoring of evidence grades and every four-state control.

The retained Go 1.25.12 and 1.26.5 darwin/arm64 `net/http` results are each:

```text
TP=39 FP=0 FN=0 TN=39
UNKNOWN=39/39 UNSUPPORTED=39/39
evidence_grade_mismatches=0 unresolved=0 unexpected=0
publication_state=FINAL coverage_verdict=COMPLETE
```

Each ordinary product was invoked without framework, mode, or entry-point
overrides. Its raw coverage is intentionally `CANNOT_CERTIFY` because the 29
UNKNOWN controls are real capability gaps. Qualification becomes `COMPLETE` only
when the observed gaps exactly equal the locked truth set.

## Run it

Build the language corpus:

```bash
cd benchmarks/benchmark-go-v1/corpus/language
GOTOOLCHAIN=go1.25.12 go build \
  -o /tmp/rig-benchmark-go-language-corpus \
  ./cmd/rig-benchmark-go-language-corpus
```

Run the ordinary application path:

```bash
cd ../../../..
rig run benchmarks/benchmark-go-v1/apps/net-http-product
```

Exercise an unqualified coordinate separately, retaining its machine error record,
then score the sealed result:

```bash
set +e
rig run benchmarks/benchmark-go-v1/controls/unqualified-coordinate \
  > /tmp/benchmark-go-unsupported-result.json 2>&1
test "$?" -eq 30
set -e

rig benchmark-go \
  --suite benchmarks/benchmark-go-v1 \
  --corpus-executable /tmp/rig-benchmark-go-language-corpus \
  --scan-result /path/to/run/go-scan-result.json \
  --unsupported-result /tmp/benchmark-go-unsupported-result.json \
  --projected-evidence-output /path/to/run/product-evidence.json \
  --output /path/to/run/benchmark-go-qualification.json
```

`make test-go-benchmark` runs the checked-in native benchmark gates. A recognized
call without the declared semantic evidence is an FN, and a finding on a safe case
is an FP.

## Promotion status

The locked target family includes Go 1.25.12 and 1.26.5 on darwin/arm64 and
linux/amd64, standard `net/http`, and Gin 1.12.0. Fresh full product runs close all
four `net/http` runtime coordinates at 39 TP / 0 FP / 0 FN / 39 TN. Gin 1.12.0 on
Go 1.26.5 closes the same full denominator at 39 TP / 0 FP / 0 FN / 39 TN, with
all evidence grades and four-state controls exact.

Promotion remains blocked until three independently curated held-out repositories
collectively close all 39 vulnerable and safe categories and include both locked
framework families. Held-out evidence is deliberately not developed in this tree;
the scorer accepts signed projections through `--held-out-evidence` and refuses
promotion otherwise. The scorer also emits one
`runtime_coordinate_evidence_missing:<coordinate>` or
`framework_coordinate_evidence_missing:<coordinate>` reason for every member of
the locked target family that is not represented by locked or held-out evidence.
This keeps an absent linux/amd64 or Gin run visible instead of hiding it behind the
generic held-out gate.
