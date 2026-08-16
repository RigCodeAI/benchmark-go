package benchmarkgo

func dataSpecs() []CategorySpec {
	return []CategorySpec{
		{"CWE-200", "DIRECT", "classified data crosses a public response boundary", "RUNTIME_VALUE_FLOW", secretExposed, secretRedacted},
		{"CWE-201", "DIRECT", "classified data crosses an outbound service boundary", "RUNTIME_VALUE_FLOW", secretExposed, secretRedacted},
		{"CWE-328", "ADAPTED", "a security operation selects a weak hash", "RUNTIME_PROPERTY", weakPrimitive, strongPrimitive},
		{"CWE-330", "ADAPTED", "a security operation selects a predictable random source", "RUNTIME_PROPERTY", weakPrimitive, strongPrimitive},
		{"CWE-532", "DIRECT", "classified data crosses a logging boundary", "RUNTIME_VALUE_FLOW", secretExposed, secretRedacted},
		{"CWE-614", "DIRECT", "a sensitive cookie lacks transport protection", "RUNTIME_PROPERTY", weakPrimitive, strongPrimitive},
	}
}
