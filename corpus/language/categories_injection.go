package benchmarkgo

func direct(category, statement, grade string) CategorySpec {
	return CategorySpec{category, "DIRECT", statement, grade, structureChanged, structureBound}
}

func adapted(category, statement, grade string) CategorySpec {
	return CategorySpec{category, "ADAPTED", statement, grade, structureChanged, structureBound}
}

func injectionSpecs() []CategorySpec {
	return []CategorySpec{
		direct("CWE-113", "request data changes HTTP header structure", "RUNTIME_SEMANTIC"),
		{"CWE-116", "DIRECT", "untrusted output remains active in its destination context", "RUNTIME_SEMANTIC", contextActive, contextEncoded},
		direct("CWE-1336", "request data changes template grammar", "RUNTIME_SEMANTIC"),
		{"CWE-22", "DIRECT", "request data escapes a filesystem root", "RUNTIME_EFFECT", pathEscaped, pathContained},
		direct("CWE-501", "less-trusted data crosses a trust boundary", "RUNTIME_VALUE_FLOW"),
		adapted("CWE-502", "request data selects executable deserialization behavior", "RUNTIME_SEMANTIC"),
		direct("CWE-601", "request data selects an untrusted redirect", "RUNTIME_SEMANTIC"),
		adapted("CWE-611", "XML input resolves an external entity", "RUNTIME_EFFECT"),
		adapted("CWE-643", "request data changes XML query grammar", "RUNTIME_SEMANTIC"),
		adapted("CWE-776", "XML input expands beyond the resource budget", "RUNTIME_EFFECT"),
		direct("CWE-78", "request data changes process invocation semantics", "RUNTIME_SEMANTIC"),
		{"CWE-79", "DIRECT", "request data remains executable in an output context", "RUNTIME_SEMANTIC", contextActive, contextEncoded},
		direct("CWE-89", "request data changes SQL grammar", "RUNTIME_SEMANTIC"),
		direct("CWE-90", "request data changes directory query grammar", "RUNTIME_SEMANTIC"),
		{"CWE-918", "DIRECT", "request data selects an outbound authority", "RUNTIME_EFFECT", destinationControlled, destinationAllowlisted},
		direct("CWE-94", "request data reaches executable code grammar", "RUNTIME_EFFECT"),
		adapted("CWE-943", "request data changes a NoSQL query operator", "RUNTIME_SEMANTIC"),
	}
}
