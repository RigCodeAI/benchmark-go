package benchmarkgo

import (
	"fmt"
	"runtime"
	"sort"
)

func CategorySpecs() []CategorySpec {
	var specs []CategorySpec
	specs = append(specs, accessSpecs()...)
	specs = append(specs, availabilitySpecs()...)
	specs = append(specs, dataSpecs()...)
	specs = append(specs, injectionSpecs()...)
	specs = append(specs, goSpecificSpecs()...)
	sort.Slice(specs, func(i, j int) bool { return specs[i].Category < specs[j].Category })
	return specs
}

func Execute() []CaseResult {
	var output []CaseResult
	for _, category := range CategorySpecs() {
		for _, control := range Controls {
			output = append(output, category.Execute(control))
		}
	}
	sort.Slice(output, func(i, j int) bool { return output[i].CaseID < output[j].CaseID })
	return output
}

func Validate() error {
	specs := CategorySpecs()
	if len(specs) != 39 {
		return fmt.Errorf("benchmark_go_category_count_mismatch")
	}
	categories := map[string]bool{}
	for _, spec := range specs {
		if categories[spec.Category] || spec.Mapping == "" || spec.SecurityStatement == "" || spec.EvidenceGrade == "" {
			return fmt.Errorf("benchmark_go_category_invalid")
		}
		categories[spec.Category] = true
	}
	cases := Execute()
	if len(cases) != 156 {
		return fmt.Errorf("benchmark_go_control_count_mismatch")
	}
	ids := map[string]bool{}
	for _, result := range cases {
		if ids[result.CaseID] || result.Expected != result.Observed || result.Witness == "" {
			return fmt.Errorf("benchmark_go_control_invalid")
		}
		ids[result.CaseID] = true
	}
	return nil
}

func RuntimeCoordinate() string {
	return "go-" + runtime.Version()[2:] + "-" + runtime.GOOS + "-" + runtime.GOARCH
}
