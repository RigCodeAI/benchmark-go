#!/bin/sh
set -eu

ROOT=$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)
cd "$ROOT"

cargo fmt --all -- --check
cargo test --locked
cargo clippy --locked --all-targets -- -D warnings
cargo run --quiet --locked -- verify

(cd corpus/language && GOTOOLCHAIN=go1.25.12 go test ./...)
(cd apps/net-http-product && GOTOOLCHAIN=go1.25.12 go test ./...)
(cd apps/gin-product && GOTOOLCHAIN=go1.26.5 go test ./...)

echo "BenchmarkGo verification complete"
