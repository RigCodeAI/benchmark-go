# Go target security coverage

Status: complete four-coordinate `net/http` and Gin 1.12.0 qualification
denominators; promotion pending independent evidence, version 5, 2026-08-16

Go applications are native binaries, but they retain the Go runtime for garbage
collection, scheduling, stacks, maps, channels, and other language services. Rig
does not require an interpreter or a general-purpose attach API. It instruments a
private source snapshot before compilation and embeds a small observation adapter
in the resulting native application.

## Product path

Ordinary `rig run /path/to/project` now owns the Go path end to end:

1. inventory `go.mod` files and main packages without importing or executing the
   repository;
2. select an exact qualified Go toolchain and fail closed when the declared
   `toolchain` directive is newer;
3. copy the immutable snapshot into a private disposable workspace, reject
   symlinks, and enforce file and byte budgets;
4. provision the content-addressed Go adapter into only that disposable module;
5. discover and rewrite qualified `net/http` routes, sources, sinks, helpers,
   aliases, async/context boundaries, and semantic operations;
6. build and launch the native target on loopback;
7. issue HMAC-scoped request capabilities and drive query, path, form, JSON,
   header, cookie, multipart, body, middleware, context, and principal sources;
8. collect authenticated handler, sink, semantic-property, fence, checkpoint,
   completion, access-control, workflow, concurrency, and service-interaction
   facts;
9. seal transcripts, reduce journeys, reconcile coverage, publish developer
   JSON/SARIF/Markdown, and authenticate final readback.

`rig targets` projects every READY Go main package into the same
`rig-application-inventory/v1` contract used by the CLI, VS Code extension, and
coding-agent hooks. Go targets use `mode: http`; they no longer depend on Python
application discovery or appear as `DETECTED_NOT_RUNNABLE` in editor inventory.
`rig feedback run` binds the native adapter and publication to the exact Git
snapshot selected by the editor or hook, then imports the authenticated Go result
through the same daemon gate used for Python. The disposable source-tree digest
remains separately retained in preparation and build-provenance artifacts.

The original repository and its module files are never modified. Traversal effect
canaries are created only inside the disposable application copy.

## Qualified local coordinate

The retained full ordinary-product results are:

- Go 1.25.12 and Go 1.26.5;
- darwin/arm64 and linux/amd64;
- standard-library `net/http`;
- `go-build-wrapper-v1`;
- 39 categories, with one vulnerable, safe, unknown, and unsupported control per
  category.

Both exact toolchains passed fresh full `net/http` product execution on both
retained host coordinates. Gin 1.12.0 on Go 1.26.5 also passed the complete
39-category ordinary-product denominator on darwin/arm64. Go's
official policy supports a release until two
newer major releases exist, which is why the qualification family contains the
currently supported 1.25 and 1.26 lines. Exact patch and host evidence is still
required before Rig grants product authority.

## Category coverage

The local `net/http` application closes these inherited categories:

- HTTP header and output encoding: CWE-113, CWE-116;
- template and code execution: CWE-1336, CWE-94;
- sensitive-data response, outbound, log, and trust-boundary flow: CWE-200,
  CWE-201, CWE-501, CWE-532;
- filesystem, command, SQL, LDAP, NoSQL, XPath, SSRF, redirect, XSS, and
  deserialization: CWE-22, CWE-78, CWE-89, CWE-90, CWE-943, CWE-643, CWE-918,
  CWE-601, CWE-79, CWE-502;
- XML external entities and expansion: CWE-611, CWE-776;
- weak hashing, randomness, cookies, and resource exhaustion: CWE-328, CWE-330,
  CWE-614, CWE-400;
- authorization, authentication, CSRF, tenant isolation, workflows, and
  multi-service behavior: CWE-284, CWE-287, CWE-306, CWE-352, CWE-639, CWE-840,
  CWE-841, CWE-862, CWE-863.

Go-specific controls cover CGO boundary contracts, goroutine leaks, unbounded HTTP
body reads, concurrent map access, and `text/template` versus `html/template`
context confusion.

Evidence is category-specific. Examples include SQL grammar versus bound data,
resolved filesystem effects, process outcome, completed outbound destination,
active HTML or template evaluation, parser configuration/effect, runtime security
properties, synchronized differential state, and authenticated downstream
interactions. Safe controls become TNs only after their exact obligations close.

## Retained full result

The fresh arbitrary-repository command was:

```bash
rig run apps/net-http-product
```

No framework, entry point, mode, or benchmark-only launch path was supplied. The
run produced 1,237 planned requests, 1,230 attempted requests, seven explicit
transport exclusions, 24 sealed transcripts, zero failed requests, acknowledged
producer completion, and authenticated final publication. Qualification projected
the ordinary result onto the locked four-state truth as:

```text
TP=39 FP=0 FN=0 TN=39
UNKNOWN=39/39
UNSUPPORTED=39/39
evidence_grade_mismatches=0
unresolved=0
unexpected_observations=0
publication_state=FINAL
coverage_verdict=COMPLETE
```

The raw product result intentionally says `CANNOT_CERTIFY`: the benchmark includes
29 explicit unsupported semantic coordinates so it can prove that gaps remain
unknown. The qualification envelope is `COMPLETE` only because those exact 29 gaps
match the independently locked UNKNOWN truth—an extra or missing gap fails closed.

Run the native corpus and scorer as documented in the
[BenchmarkGo README](../README.md).

The retained Gin command was likewise the ordinary arbitrary-repository path:

```bash
rig run apps/gin-product
```

It produced 1,237 planned probes, 1,230 attempted requests, seven exact transport
exclusions, 24 sealed transcripts, zero failed requests, and acknowledged
producer completion. Its locked qualification result is also 39 TP / 0 FP /
0 FN / 39 TN, with all UNKNOWN and UNSUPPORTED controls correct and zero
evidence-grade mismatches.

## Remaining promotion gates

The local denominator is closed, but broad Go support is not yet qualified:

- three genuinely independent held-out repositories must collectively cover all
  39 vulnerable and safe categories; evidence authored during adapter development
  cannot satisfy this gate;
- Gin 1.12.0 closes discovery, instrumentation, launch, runtime collection, and
  the complete 39-category denominator. Its independently curated held-out
  result remains open. Chi, Echo, Fiber, gRPC, and custom routers remain
  unsupported;
- provider breadth must expand beyond the exact standard-library and modeled
  coordinates in the benchmark (database drivers/ORMs, HTTP clients, parsers,
  templates, telemetry, archives/uploads, DNS/TLS/proxies, and native libraries);
- generated code, build tags, workspaces, vendoring, module replacements, CGO,
  race builds, cross-compilation, and custom build tools require exact policy and
  evidence for each affected application;
- unresolved frameworks, routes, sources, sinks, goroutine causality, semantic
  coordinates, or build behavior remain `UNKNOWN` and cannot certify a clean scan.

The machine promotion gate reports independent held-out evidence separately from
each missing runtime and framework coordinate. The broader framework/host items
above are therefore executable gates, not documentation-only caveats, and remain
required before the support policy can claim those coordinates to customers.
