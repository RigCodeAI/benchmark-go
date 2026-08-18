.PHONY: run run-gin score verify catalog

run:
	./runBenchmark.sh net-http

run-gin:
	./runBenchmark.sh gin

score:
	@test -n "$(RESULTS)" || { echo "usage: make score RESULTS=path/to/results.sarif" >&2; exit 2; }
	./scoreBenchmark.sh --results "$(RESULTS)" $(if $(OUTPUT),--output-dir "$(OUTPUT)",)

verify:
	./verifyBenchmark.sh

catalog:
	cargo run --quiet --locked -- catalog
