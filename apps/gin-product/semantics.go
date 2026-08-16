package main

import (
	"context"
	"crypto/md5" // #nosec G501 -- intentionally vulnerable benchmark control.
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io"
	"math/big"
	mathrand "math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	texttemplate "text/template"
)

type semanticControl string

const (
	vulnerableControl semanticControl = "vulnerable"
	safeControl       semanticControl = "safe"
	unknownControl    semanticControl = "unknown"
)

func registerSemanticRoutes(mux *http.ServeMux) {
	registerSemanticCategory(mux, "cwe-113", "GET /qualification/cwe-113/vulnerable", "GET /qualification/cwe-113/safe", "GET /qualification/cwe-113/unknown")
	registerSemanticCategory(mux, "cwe-116", "GET /qualification/cwe-116/vulnerable", "GET /qualification/cwe-116/safe", "GET /qualification/cwe-116/unknown")
	registerSemanticCategory(mux, "cwe-1336", "GET /qualification/cwe-1336/vulnerable", "GET /qualification/cwe-1336/safe", "GET /qualification/cwe-1336/unknown")
	registerSemanticCategory(mux, "cwe-200", "GET /qualification/cwe-200/vulnerable", "GET /qualification/cwe-200/safe", "GET /qualification/cwe-200/unknown")
	registerSemanticCategory(mux, "cwe-201", "GET /qualification/cwe-201/vulnerable", "GET /qualification/cwe-201/safe", "GET /qualification/cwe-201/unknown")
	registerSemanticCategory(mux, "cwe-22", "GET /qualification/cwe-22/vulnerable", "GET /qualification/cwe-22/safe", "GET /qualification/cwe-22/unknown")
	registerSemanticCategory(mux, "cwe-328", "GET /qualification/cwe-328/vulnerable", "GET /qualification/cwe-328/safe", "GET /qualification/cwe-328/unknown")
	registerSemanticCategory(mux, "cwe-330", "GET /qualification/cwe-330/vulnerable", "GET /qualification/cwe-330/safe", "GET /qualification/cwe-330/unknown")
	registerSemanticCategory(mux, "cwe-400", "GET /qualification/cwe-400/vulnerable", "GET /qualification/cwe-400/safe", "GET /qualification/cwe-400/unknown")
	registerSemanticCategory(mux, "cwe-501", "GET /qualification/cwe-501/vulnerable", "GET /qualification/cwe-501/safe", "GET /qualification/cwe-501/unknown")
	registerSemanticCategory(mux, "cwe-502", "GET /qualification/cwe-502/vulnerable", "GET /qualification/cwe-502/safe", "GET /qualification/cwe-502/unknown")
	registerSemanticCategory(mux, "cwe-532", "GET /qualification/cwe-532/vulnerable", "GET /qualification/cwe-532/safe", "GET /qualification/cwe-532/unknown")
	registerSemanticCategory(mux, "cwe-601", "GET /qualification/cwe-601/vulnerable", "GET /qualification/cwe-601/safe", "GET /qualification/cwe-601/unknown")
	registerSemanticCategory(mux, "cwe-611", "GET /qualification/cwe-611/vulnerable", "GET /qualification/cwe-611/safe", "GET /qualification/cwe-611/unknown")
	registerSemanticCategory(mux, "cwe-614", "GET /qualification/cwe-614/vulnerable", "GET /qualification/cwe-614/safe", "GET /qualification/cwe-614/unknown")
	registerSemanticCategory(mux, "cwe-643", "GET /qualification/cwe-643/vulnerable", "GET /qualification/cwe-643/safe", "GET /qualification/cwe-643/unknown")
	registerSemanticCategory(mux, "cwe-776", "GET /qualification/cwe-776/vulnerable", "GET /qualification/cwe-776/safe", "GET /qualification/cwe-776/unknown")
	registerSemanticCategory(mux, "cwe-78", "GET /qualification/cwe-78/vulnerable", "GET /qualification/cwe-78/safe", "GET /qualification/cwe-78/unknown")
	registerSemanticCategory(mux, "cwe-79", "GET /qualification/cwe-79/vulnerable", "GET /qualification/cwe-79/safe", "GET /qualification/cwe-79/unknown")
	registerSemanticCategory(mux, "cwe-89", "GET /qualification/cwe-89/vulnerable", "GET /qualification/cwe-89/safe", "GET /qualification/cwe-89/unknown")
	registerSemanticCategory(mux, "cwe-90", "GET /qualification/cwe-90/vulnerable", "GET /qualification/cwe-90/safe", "GET /qualification/cwe-90/unknown")
	registerSemanticCategory(mux, "cwe-918", "GET /qualification/cwe-918/vulnerable", "GET /qualification/cwe-918/safe", "GET /qualification/cwe-918/unknown")
	registerSemanticCategory(mux, "cwe-94", "GET /qualification/cwe-94/vulnerable", "GET /qualification/cwe-94/safe", "GET /qualification/cwe-94/unknown")
	registerSemanticCategory(mux, "cwe-943", "GET /qualification/cwe-943/vulnerable", "GET /qualification/cwe-943/safe", "GET /qualification/cwe-943/unknown")
	registerSemanticCategory(mux, "go-cgo-boundary", "GET /qualification/go-cgo-boundary/vulnerable", "GET /qualification/go-cgo-boundary/safe", "GET /qualification/go-cgo-boundary/unknown")
	registerSemanticCategory(mux, "go-goroutine-leak", "GET /qualification/go-goroutine-leak/vulnerable", "GET /qualification/go-goroutine-leak/safe", "GET /qualification/go-goroutine-leak/unknown")
	registerSemanticCategory(mux, "go-http-body-limit", "GET /qualification/go-http-body-limit/vulnerable", "GET /qualification/go-http-body-limit/safe", "GET /qualification/go-http-body-limit/unknown")
	registerSemanticCategory(mux, "go-map-concurrent-access", "GET /qualification/go-map-concurrent-access/vulnerable", "GET /qualification/go-map-concurrent-access/safe", "GET /qualification/go-map-concurrent-access/unknown")
	registerSemanticCategory(mux, "go-template-context-confusion", "GET /qualification/go-template-context-confusion/vulnerable", "GET /qualification/go-template-context-confusion/safe", "GET /qualification/go-template-context-confusion/unknown")
}

func registerSemanticCategory(mux *http.ServeMux, category, vulnerable, safe, unknown string) {
	// The literal patterns above are the controller's immutable route inventory.
	// Route registration is deliberately centralized so framework discovery must
	// resolve helper arguments rather than accept only direct global registrations.
	mux.HandleFunc(vulnerable, semanticHandler(category, vulnerableControl))
	mux.HandleFunc(safe, semanticHandler(category, safeControl))
	mux.HandleFunc(unknown, semanticHandler(category, unknownControl))
}

func semanticHandler(category string, control semanticControl) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		value := request.URL.Query().Get("secret")
		switch control {
		case unknownControl:
			opaqueSemanticBoundary(category, value)
			writer.WriteHeader(http.StatusNoContent)
		case vulnerableControl:
			vulnerableSemantic(writer, request, category, value)
		case safeControl:
			safeSemantic(writer, request, category, value)
		}
	}
}

// HTTP body-limit qualification needs the request body itself to be both the
// runtime source and the operation observed by io.ReadAll. Keeping this as a
// dedicated handler prevents a synthetic query value from standing in for the
// concrete body whose boundedness is being decided.
func httpBodyLimitHandler(bounded bool) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if bounded {
			_, _ = io.ReadAll(http.MaxBytesReader(writer, request.Body, 64<<10))
		} else {
			_, _ = io.ReadAll(request.Body)
		}
		writer.WriteHeader(http.StatusNoContent)
	}
}

func vulnerableSemantic(writer http.ResponseWriter, request *http.Request, category, value string) {
	switch category {
	case "cwe-113":
		writer.Header().Set("X-Rig-Value", value)
	case "cwe-116":
		_, _ = writer.Write([]byte(template.HTML(value))) // #nosec G203 -- intentionally unsafe output control.
	case "cwe-79":
		_, _ = fmt.Fprint(writer, value)
	case "go-template-context-confusion":
		parsed, _ := texttemplate.New("text-context").Parse(value)
		if parsed != nil {
			_ = parsed.Execute(writer, value)
		}
	case "cwe-1336":
		parsed, _ := texttemplate.New("dynamic").Parse(value)
		if parsed != nil {
			_ = parsed.Execute(writer, value)
		}
	case "cwe-200":
		_, _ = fmt.Fprintf(writer, "secret=%s", value)
	case "cwe-201":
		_, _ = http.NewRequestWithContext(request.Context(), http.MethodPost, "http://127.0.0.1:1/collect", strings.NewReader(value))
	case "cwe-22":
		_, _ = os.ReadFile(value)
	case "cwe-328":
		_ = md5.Sum([]byte(value)) // #nosec G401 -- vulnerable control.
	case "cwe-330":
		_ = mathrand.Int63()
	case "cwe-400":
		_ = resourceOperation(value, false)
	case "cwe-501":
		_ = trustBoundaryOperation(value, false)
	case "cwe-502":
		var unrestrictedDocument effectfulDocument
		_ = json.Unmarshal([]byte(value), &unrestrictedDocument)
	case "cwe-532":
		fmt.Printf("benchmark secret=%s\n", value)
	case "cwe-601":
		http.Redirect(writer, request, value, http.StatusFound)
		return
	case "cwe-611", "cwe-776":
		if category == "cwe-611" {
			_ = xmlExternalOperation(value, false)
		} else {
			_ = xmlExpansionOperation(value, false)
		}
	case "cwe-614":
		http.SetCookie(writer, &http.Cookie{Name: "session", Value: value})
	case "cwe-643":
		_ = xpathOperation(value, false)
	case "cwe-78":
		_ = exec.Command("/bin/sh", "-c", value).Run() // #nosec G204 -- vulnerable control.
	case "cwe-89":
		statement := "SELECT * FROM item WHERE name='" + value + "'"
		_, _ = semanticDatabase.ExecContext(request.Context(), statement)
	case "cwe-90":
		_ = ldapOperation(value, false)
	case "cwe-918":
		response, _ := http.Get(value) // #nosec G107 -- vulnerable control.
		if response != nil {
			_ = response.Body.Close()
		}
	case "cwe-94":
		parsed, _ := texttemplate.New("code").Parse(value)
		if parsed != nil {
			_ = parsed.Execute(io.Discard, value)
		}
	case "cwe-943":
		_ = nosqlOperation(value, false)
	case "go-cgo-boundary":
		_ = cgoBoundaryOperation(value, false)
	case "go-goroutine-leak":
		_ = goroutineOperation(value, false)
	case "go-map-concurrent-access":
		_ = mapConcurrencyOperation(value, false)
	}
	writer.WriteHeader(http.StatusNoContent)
}

func safeSemantic(writer http.ResponseWriter, request *http.Request, category, value string) {
	switch category {
	case "cwe-113":
		writer.Header().Set("X-Rig-Value", strings.NewReplacer("\r", "", "\n", "").Replace(value))
	case "cwe-116":
		_, _ = writer.Write([]byte(html.EscapeString(value)))
	case "cwe-79":
		encoded := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;", "'", "&#39;").Replace(value)
		_, _ = fmt.Fprint(writer, encoded)
	case "cwe-1336":
		parsed := template.Must(template.New("fixed").Parse("{{.}}"))
		_ = parsed.Execute(writer, value)
	case "go-template-context-confusion":
		parsed := template.Must(template.New("fixed-context").Parse("{{.}}"))
		_ = parsed.Execute(writer, value)
	case "cwe-200":
		_, _ = fmt.Fprintf(writer, "%s", "[REDACTED]")
	case "cwe-201":
		_, _ = http.NewRequestWithContext(request.Context(), http.MethodPost, "http://127.0.0.1:1/collect", strings.NewReader("[REDACTED]"))
	case "cwe-532":
		_, _ = fmt.Printf("%s", "[REDACTED]")
	case "cwe-22":
		_, _ = os.ReadFile(filepath.Join(os.TempDir(), "rig-go-safe-file"))
	case "cwe-328":
		_ = sha256.Sum256([]byte(value))
	case "cwe-330":
		_, _ = rand.Int(rand.Reader, big.NewInt(1<<30))
	case "cwe-400":
		_ = resourceOperation(value, true)
	case "cwe-501":
		_ = trustBoundaryOperation(value, true)
	case "cwe-502":
		var restrictedDocument map[string]string
		_ = json.Unmarshal([]byte(`{"value":"safe"}`), &restrictedDocument)
	case "cwe-601":
		http.Redirect(writer, request, "/safe", http.StatusFound)
		return
	case "cwe-611", "cwe-776":
		if category == "cwe-611" {
			_ = xmlExternalOperation(value, true)
		} else {
			_ = xmlExpansionOperation(value, true)
		}
	case "cwe-614":
		http.SetCookie(writer, &http.Cookie{Name: "session", Value: value, Secure: true, HttpOnly: true})
	case "cwe-643":
		_ = xpathOperation(value, true)
	case "cwe-78":
		_ = exec.Command("/usr/bin/printf", "%s", value).Run()
	case "cwe-89":
		_, _ = semanticDatabase.ExecContext(request.Context(), "SELECT * FROM item WHERE name=?", value)
	case "cwe-90":
		_ = ldapOperation(value, true)
	case "cwe-918":
		response, _ := http.Get("http://127.0.0.1:1/fixed")
		if response != nil {
			_ = response.Body.Close()
		}
	case "cwe-94":
		parsed := template.Must(template.New("fixed-code").Parse("{{.}}"))
		_ = parsed.Execute(writer, value)
	case "cwe-943":
		_ = nosqlOperation(value, true)
	case "go-cgo-boundary":
		_ = cgoBoundaryOperation(value, true)
	case "go-goroutine-leak":
		_ = goroutineOperation(value, true)
	case "go-map-concurrent-access":
		_ = mapConcurrencyOperation(value, true)
	}
	writer.WriteHeader(http.StatusNoContent)
}

var trustStore = map[string]string{}

var semanticDatabase semanticSQLExecutor

type semanticSQLExecutor struct{}

func (semanticSQLExecutor) ExecContext(_ context.Context, _ string, _ ...any) (sql.Result, error) {
	return semanticSQLResult(0), nil
}

type semanticSQLResult int64

func (result semanticSQLResult) LastInsertId() (int64, error) { return int64(result), nil }
func (result semanticSQLResult) RowsAffected() (int64, error) { return int64(result), nil }

type effectfulDocument struct{}

func (*effectfulDocument) UnmarshalJSON(document []byte) error {
	return os.WriteFile("rig-owned-go-deserialization-effect", document, 0o600)
}

func resourceOperation(value string, bounded bool) string {
	iterations := len(value) * 1024
	if bounded && iterations > 1024 {
		iterations = 1024
	}
	for index := 0; index < iterations; index++ {
		_ = sha256.Sum256([]byte(value))
	}
	if bounded {
		return "bounded"
	}
	return "input_scaled_work"
}

func trustBoundaryOperation(value string, protected bool) string {
	if protected {
		trustStore["validated"] = "safe"
		return "protected"
	}
	trustStore["unvalidated"] = value
	return "unprotected"
}

func xmlExternalOperation(value string, protected bool) string {
	if protected {
		return "disabled"
	}
	return strings.ReplaceAll(value, "&rig;", "RIG_EXTERNAL_EFFECT")
}

func xmlExpansionOperation(value string, protected bool) string {
	if protected {
		return "disabled"
	}
	return strings.Repeat(value, 4)
}

func xpathOperation(value string, protected bool) string {
	if protected {
		value = strings.NewReplacer("'", "&apos;", "\"", "&quot;").Replace(value)
		return "escaped_literal://item[text()='" + value + "']"
	}
	return "query_grammar://item[text()='" + value + "']"
}

func ldapOperation(value string, protected bool) string {
	if protected {
		value = strings.NewReplacer("\\", "\\5c", "*", "\\2a", "(", "\\28", ")", "\\29", "\x00", "\\00").Replace(value)
		return "escaped_assertion_value:(&(objectClass=person)(uid=" + value + "))"
	}
	return "filter_grammar:(&(objectClass=person)(uid=" + value + "))"
}

func nosqlOperation(value string, protected bool) string {
	if protected {
		return `literal_value:{"name":` + value + `}`
	}
	return `operator_document:{"$where":` + value + `}`
}

func goroutineOperation(value string, bounded bool) string {
	if !bounded {
		go func() { select {} }()
		return "leaked"
	}
	done := make(chan struct{})
	go func() { defer close(done); _ = value }()
	<-done
	return "joined"
}

func mapConcurrencyOperation(value string, protected bool) string {
	if protected {
		shared := map[string]string{}
		var lock sync.Mutex
		var wait sync.WaitGroup
		for index := 0; index < 2; index++ {
			wait.Add(1)
			go func() {
				defer wait.Done()
				lock.Lock()
				defer lock.Unlock()
				shared[value] = value
			}()
		}
		wait.Wait()
		return "serialized"
	}
	var active atomic.Int32
	var overlap atomic.Bool
	var wait sync.WaitGroup
	ready := make(chan struct{}, 2)
	release := make(chan struct{})
	for index := 0; index < 2; index++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			if active.Add(1) > 1 {
				overlap.Store(true)
			}
			ready <- struct{}{}
			<-release
			runtime.Gosched()
			active.Add(-1)
		}()
	}
	<-ready
	<-ready
	close(release)
	wait.Wait()
	if overlap.Load() {
		return "overlap_observed"
	}
	return "overlap_not_observed"
}

func opaqueSemanticBoundary(category, value string) { _, _ = category, value }
