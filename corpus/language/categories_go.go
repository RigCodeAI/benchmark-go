package benchmarkgo

func goSpecificSpecs() []CategorySpec {
	return []CategorySpec{
		{"GO-CGO-BOUNDARY", "ADAPTED", "request data reaches a cgo boundary without a checked contract", "RUNTIME_MEMORY", memoryUnchecked, memoryChecked},
		{"GO-GOROUTINE-LEAK", "ADAPTED", "request-controlled goroutines outlive their bounded lifecycle", "RUNTIME_EFFECT", unboundedEffect, boundedEffect},
		{"GO-HTTP-BODY-LIMIT", "ADAPTED", "request bodies are consumed without an enforced size limit", "RUNTIME_PROPERTY", unboundedEffect, boundedEffect},
		{"GO-MAP-CONCURRENT-ACCESS", "ADAPTED", "concurrent handlers access a map without synchronization", "RUNTIME_DIFFERENTIAL", policyBypassed, policyEnforced},
		{"GO-TEMPLATE-CONTEXT-CONFUSION", "ADAPTED", "template data is emitted in the wrong escaping context", "RUNTIME_SEMANTIC", contextActive, contextEncoded},
	}
}
