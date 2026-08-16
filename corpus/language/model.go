package benchmarkgo

import "strings"

type Control string

const (
	Vulnerable  Control = "vulnerable"
	Safe        Control = "safe"
	Unknown     Control = "unknown"
	Unsupported Control = "unsupported"
)

var Controls = [...]Control{Vulnerable, Safe, Unknown, Unsupported}

type Disposition string

const (
	Finding            Disposition = "FINDING"
	Clean              Disposition = "CLEAN"
	CapabilityGap      Disposition = "UNKNOWN"
	UnsupportedRuntime Disposition = "UNSUPPORTED"
)

type ProbeResult struct {
	Disposition Disposition
	Witness     string
}

type Probe func() ProbeResult

type CategorySpec struct {
	Category          string
	Mapping           string
	SecurityStatement string
	EvidenceGrade     string
	Vulnerable        Probe
	Safe              Probe
}

type CaseResult struct {
	CaseID   string
	Category string
	Control  Control
	Expected Disposition
	Observed Disposition
	Witness  string
}

func (spec CategorySpec) Execute(control Control) CaseResult {
	probe := ProbeResult{}
	switch control {
	case Vulnerable:
		probe = spec.Vulnerable()
	case Safe:
		probe = spec.Safe()
	case Unknown:
		probe = ProbeResult{CapabilityGap, "required semantic fact is intentionally opaque"}
	case Unsupported:
		probe = ProbeResult{UnsupportedRuntime, "required provider coordinate is intentionally unqualified"}
	}
	return CaseResult{
		CaseID:   strings.ToLower(spec.Category) + "-" + string(control),
		Category: spec.Category,
		Control:  control,
		Expected: expected(control),
		Observed: probe.Disposition,
		Witness:  probe.Witness,
	}
}

func expected(control Control) Disposition {
	switch control {
	case Vulnerable:
		return Finding
	case Safe:
		return Clean
	case Unknown:
		return CapabilityGap
	default:
		return UnsupportedRuntime
	}
}
