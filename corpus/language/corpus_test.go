package benchmarkgo

import "testing"

func TestAllCategoriesHaveExecutableFourWayControls(t *testing.T) {
	if err := Validate(); err != nil {
		t.Fatal(err)
	}
	counts := map[Disposition]int{}
	for _, result := range Execute() {
		counts[result.Observed]++
	}
	for _, disposition := range []Disposition{Finding, Clean, CapabilityGap, UnsupportedRuntime} {
		if counts[disposition] != 39 {
			t.Fatalf("%s=%d", disposition, counts[disposition])
		}
	}
}
