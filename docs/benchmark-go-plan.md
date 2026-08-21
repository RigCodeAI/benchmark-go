# BenchmarkGo qualification plan

Status: four-coordinate `net/http` and Gin 1.12.0 denominators complete;
independent promotion gates open, version 5, 2026-08-16

## Completed

BenchmarkGo now has a closed, native qualification system:

- 39 categories and 156 vulnerable/safe/unknown/unsupported controls;
- a five-second native Go language corpus;
- 87 ordinary product cases, 30 controller cases, 11 source contracts, and 27
  foundation/hostile/coordinate cases;
- exact truth, schemas, source digests, deterministic replay, and source budgets;
- an ordinary `sivere run` `net/http` application covering request sources, classic
  injection, semantic properties, sensitive data, Go-specific behavior, and
  stateful controller protocols;
- native evidence projection and scoring with authenticated/sealed publication.

Fresh Go 1.25.12 and 1.26.5 runs on both darwin/arm64 and linux/amd64 each score
39 TP / 0 FP / 0 FN / 39 TN, pass all 39 UNKNOWN and 39 UNSUPPORTED controls,
have zero evidence-grade mismatches, and publish `FINAL` with a `COMPLETE`
qualification envelope. The retained Linux runs execute real x86_64 binaries,
CGO, traffic, collection, reduction, and publication inside Linux; they are not
cross-compilation-only evidence.

Gin 1.12.0 on Go 1.26.5 is now a runnable exact coordinate. Its ordinary product
run covers inventory, route groups and path normalization, request-source
recognition, reflection-bounded handler correlation, loopback launch, all 39
category evidence contracts, sealed publication, and `FINAL`/`COMPLETE`. It
scores 39 TP / 0 FP / 0 FN / 39 TN with exact UNKNOWN and UNSUPPORTED controls.

## Why promotion remains closed

The benchmark intentionally separates implementation evidence from independent
generalization evidence. Promotion still requires:

1. three held-out repositories not used to develop the adapter or semantic models;
2. held-out vulnerable and safe coverage for every category;
3. both `net/http` and Gin framework evidence;
4. zero FP/FN, exact evidence grades, deterministic replay, sealed transcripts,
   authenticated readback, and `FINAL`/`COMPLETE` throughout.

The scorer enforces those requirements. For a supplied local qualification run it
reports `held_out_application_evidence_missing` plus an exact reason for every
unrepresented runtime and framework coordinate. The held-out gate cannot honestly
be closed by fixtures authored during this implementation.

## Next engineering slices

The qualification report enumerates each missing runtime and framework coordinate
as a machine-readable reason code. A passing darwin/arm64 `net/http` run therefore
cannot make the family look qualified while linux/amd64 or Gin evidence is absent.

1. Curate three external held-out repositories under independent truth ownership
   and run them only after the models are frozen.
2. Expand exact provider packs for widely used databases/ORMs, outbound clients,
   templates, XML/query parsers, logs/telemetry, archives/uploads, cloud SDKs, and
   native/CGO boundaries.
3. Qualify additional maintained routers in usage order: Chi, Echo, Fiber, then
   gRPC and common serverless/container launch layouts.
4. Group adjacent toolchain/framework versions only when their discovery,
   instrumentation, runtime, and semantic compatibility signatures are proven
   identical.

Unknown build behavior, routes, sources, sinks, causal goroutines, or provider
coordinates always stays `UNKNOWN`; it never contributes to a clean result.
