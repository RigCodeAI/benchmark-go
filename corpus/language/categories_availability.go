package benchmarkgo

func availabilitySpecs() []CategorySpec {
	return []CategorySpec{
		{"CWE-400", "ADAPTED", "request-controlled work exceeds its declared budget", "RUNTIME_EFFECT", unboundedEffect, boundedEffect},
	}
}
