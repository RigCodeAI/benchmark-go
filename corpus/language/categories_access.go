package benchmarkgo

func access(category, statement string) CategorySpec {
	return CategorySpec{category, "CONTROLLER", statement, "RUNTIME_DIFFERENTIAL", policyBypassed, policyEnforced}
}

func accessSpecs() []CategorySpec {
	return []CategorySpec{
		access("CWE-284", "a protected operation enforces its declared access policy"),
		access("CWE-287", "a protected operation authenticates the principal"),
		access("CWE-306", "critical functionality requires authentication"),
		access("CWE-352", "a state-changing request enforces anti-CSRF policy"),
		access("CWE-362", "concurrent requests preserve the declared invariant"),
		access("CWE-639", "a principal cannot select another tenant's object"),
		access("CWE-840", "a workflow enforces its business invariant"),
		access("CWE-841", "workflow steps cannot execute out of order"),
		access("CWE-862", "authorization is present on the protected action"),
		access("CWE-863", "authorization is correct for the tested actor and resource"),
	}
}
