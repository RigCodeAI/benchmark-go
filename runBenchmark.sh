#!/bin/sh
set -eu

ROOT=$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)
APPLICATION=${1:-net-http}

if [ "$APPLICATION" = "--help" ] || [ "$APPLICATION" = "-h" ]; then
  echo "usage: $0 [net-http|gin]"
  echo "  net-http  Go 1.25.12, http://127.0.0.1:8080 (default)"
  echo "  gin       Go 1.26.5 / Gin 1.12.0, http://127.0.0.1:8080"
  exit 0
fi

command -v go >/dev/null 2>&1 || {
  echo "Go is required. Install Go with GOTOOLCHAIN support." >&2
  exit 2
}

case "$APPLICATION" in
  net-http)
    DIRECTORY=apps/net-http-product
    DEFAULT_TOOLCHAIN=go1.25.12
    ;;
  gin)
    DIRECTORY=apps/gin-product
    DEFAULT_TOOLCHAIN=go1.26.5
    ;;
  *)
    echo "usage: $0 [net-http|gin]" >&2
    exit 2
    ;;
esac

TOOLCHAIN=${BENCHMARK_GO_TOOLCHAIN:-$DEFAULT_TOOLCHAIN}
cd "$ROOT/$DIRECTORY"
exec env GOTOOLCHAIN="$TOOLCHAIN" go run .
