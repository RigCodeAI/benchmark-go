package main

import (
	"fmt"
	"os"

	benchmarkgo "benchmark.invalid/rig-benchmark-go-language-corpus"
)

func main() {
	if err := benchmarkgo.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("runtime=%s\n", benchmarkgo.RuntimeCoordinate())
	fmt.Println("BenchmarkGo language corpus: categories=39 controls=156 vulnerable=39 safe=39 unknown=39 unsupported=39")
}
