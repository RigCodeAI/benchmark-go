package benchmarkgo

import (
	"html"
	"net/url"
	"path/filepath"
	"strings"
)

func finding(witness string) ProbeResult { return ProbeResult{Finding, witness} }
func clean(witness string) ProbeResult   { return ProbeResult{Clean, witness} }

func structureChanged() ProbeResult {
	input := "' OR 1=1 --"
	statement := "SELECT * FROM item WHERE name = '" + input + "'"
	if strings.Contains(statement, " OR 1=1 ") {
		return finding("request bytes changed an interpreter grammar")
	}
	return clean("request bytes remained data")
}

func structureBound() ProbeResult {
	statement := "SELECT * FROM item WHERE name = ?"
	parameter := "' OR 1=1 --"
	if strings.Contains(statement, parameter) {
		return finding("bound value entered interpreter grammar")
	}
	return clean("bound value remained outside interpreter grammar")
}

func contextActive() ProbeResult {
	input := "<script>globalThis.rig=1</script>"
	if strings.Contains("<main>"+input+"</main>", "<script>") {
		return finding("request bytes remained active in the output context")
	}
	return clean("request bytes were inert")
}

func contextEncoded() ProbeResult {
	body := html.EscapeString("<script>globalThis.rig=1</script>")
	if strings.Contains(body, "<script>") {
		return finding("encoded bytes remained active")
	}
	return clean("context-sensitive encoding made request bytes inert")
}

func pathEscaped() ProbeResult {
	root := "/srv/app/data"
	resolved := filepath.Clean(filepath.Join(root, "../../etc/passwd"))
	if !strings.HasPrefix(resolved, root+string(filepath.Separator)) {
		return finding("request path escaped the declared containment root")
	}
	return clean("request path remained contained")
}

func pathContained() ProbeResult {
	root := "/srv/app/data"
	resolved := filepath.Clean(filepath.Join(root, "profile.txt"))
	if !strings.HasPrefix(resolved, root+string(filepath.Separator)) {
		return finding("normalized path escaped containment")
	}
	return clean("normalized path remained under the declared root")
}

func destinationControlled() ProbeResult {
	trusted, _ := url.Parse("https://api.example.invalid")
	selected, _ := url.Parse("http://169.254.169.254/latest/meta-data")
	if selected.Host != trusted.Host {
		return finding("request input selected the outbound authority")
	}
	return clean("outbound authority was fixed by policy")
}

func destinationAllowlisted() ProbeResult {
	selected := "api.example.invalid"
	if selected == "api.example.invalid" {
		return clean("outbound authority matched the exact allowlist")
	}
	return finding("outbound authority bypassed the allowlist")
}

func policyBypassed() ProbeResult {
	if "member" != "admin" {
		return finding("a lower-privilege principal completed the protected action")
	}
	return clean("principal policy was enforced")
}

func policyEnforced() ProbeResult {
	if "member" == "admin" {
		return finding("invalid control fixture")
	}
	return clean("the lower-privilege principal was denied")
}

func secretExposed() ProbeResult {
	secret := "token_rig_secret"
	if strings.Contains("debug="+secret, secret) {
		return finding("a classified secret crossed a public boundary")
	}
	return clean("classified data was removed")
}

func secretRedacted() ProbeResult {
	if strings.Contains("debug=[REDACTED]", "token_rig_secret") {
		return finding("redaction failed")
	}
	return clean("classified data was redacted before the boundary")
}

func weakPrimitive() ProbeResult   { return finding("a qualified weak primitive was selected") }
func strongPrimitive() ProbeResult { return clean("a qualified strong primitive was selected") }
func unboundedEffect() ProbeResult { return finding("request-controlled work exceeded its budget") }
func boundedEffect() ProbeResult   { return clean("request-controlled work was capped") }
func memoryUnchecked() ProbeResult { return finding("input reached an unchecked native boundary") }
func memoryChecked() ProbeResult {
	return clean("bounds and ownership checks preceded the native boundary")
}
