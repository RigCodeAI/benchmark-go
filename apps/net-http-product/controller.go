package main

import (
	"net/http"
	"os"
	"sync"
	"time"
)

var storeMutex sync.Mutex
var raceClaims int
var workflowExecutions int

func registerControllerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /controller/login", login)
	mux.HandleFunc("GET /controller/csrf-token", csrfToken)
	mux.HandleFunc("GET /controller/cwe-284/vulnerable", allow)
	mux.HandleFunc("GET /controller/cwe-284/safe", allowIdentity("alice"))
	mux.HandleFunc("GET /controller/cwe-287/vulnerable", allow)
	mux.HandleFunc("GET /controller/cwe-287/safe", allowIdentity("alice"))
	mux.HandleFunc("GET /controller/cwe-306/vulnerable", allow)
	mux.HandleFunc("GET /controller/cwe-306/safe", allowIdentity("alice"))
	mux.HandleFunc("GET /controller/cwe-639/vulnerable", allow)
	mux.HandleFunc("GET /controller/cwe-639/safe", allowIdentity("alice"))
	mux.HandleFunc("GET /controller/cwe-862/vulnerable", allow)
	mux.HandleFunc("GET /controller/cwe-862/safe", allowIdentity("alice"))
	mux.HandleFunc("GET /controller/cwe-863/vulnerable", allow)
	mux.HandleFunc("GET /controller/cwe-863/safe", allowIdentity("admin"))
	mux.HandleFunc("POST /controller/cwe-352/vulnerable", csrfVulnerable)
	mux.HandleFunc("POST /controller/cwe-352/safe", csrfSafe)
	registerWorkflowRoutes(mux)
	registerRaceRoutes(mux)
	registerServiceRoutes(mux)
}

func login(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{Name: "rig_session", Value: request.FormValue("username"), HttpOnly: true})
	writer.WriteHeader(http.StatusOK)
}

func csrfToken(writer http.ResponseWriter, _ *http.Request) {
	http.SetCookie(writer, &http.Cookie{Name: "rig_csrf", Value: "rig-token"})
	writer.WriteHeader(http.StatusOK)
}

func csrfVulnerable(writer http.ResponseWriter, _ *http.Request) { writer.WriteHeader(http.StatusOK) }

func csrfSafe(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-CSRF-Token") != "rig-token" {
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func allow(writer http.ResponseWriter, _ *http.Request) { writer.WriteHeader(http.StatusOK) }

func allowIdentity(identity string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("rig_session")
		if err != nil || cookie.Value != identity {
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func registerWorkflowRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /controller/workflow/control/reset", workflowReset)
	mux.HandleFunc("POST /controller/workflow/control/execute", workflowLimited)
	mux.HandleFunc("POST /controller/workflow/safe/reset", workflowReset)
	mux.HandleFunc("POST /controller/workflow/safe/execute", workflowLimited)
	mux.HandleFunc("POST /controller/workflow/vulnerable/reset", workflowReset)
	mux.HandleFunc("POST /controller/workflow/vulnerable/execute", workflowUnlimited)
}

func workflowReset(writer http.ResponseWriter, _ *http.Request) {
	storeMutex.Lock()
	workflowExecutions = 0
	storeMutex.Unlock()
	writer.WriteHeader(http.StatusOK)
}

func workflowLimited(writer http.ResponseWriter, _ *http.Request)   { executeWorkflow(writer, true) }
func workflowUnlimited(writer http.ResponseWriter, _ *http.Request) { executeWorkflow(writer, false) }

func executeWorkflow(writer http.ResponseWriter, limit bool) {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	workflowExecutions++
	if limit && workflowExecutions > 1 {
		writer.WriteHeader(http.StatusConflict)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func registerRaceRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /controller/race/safe/reset", raceReset)
	mux.HandleFunc("POST /controller/race/vulnerable/reset", raceReset)
	mux.HandleFunc("POST /controller/race/safe/claim", raceSafeClaim)
	mux.HandleFunc("POST /controller/race/vulnerable/claim", raceVulnerableClaim)
	mux.HandleFunc("GET /controller/race/safe/oracle", raceSafeOracle)
	mux.HandleFunc("GET /controller/race/vulnerable/oracle", raceVulnerableOracle)
}

func raceReset(writer http.ResponseWriter, _ *http.Request) {
	storeMutex.Lock()
	raceClaims = 0
	storeMutex.Unlock()
	writer.WriteHeader(http.StatusNoContent)
}

func raceSafeClaim(writer http.ResponseWriter, _ *http.Request) {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	if raceClaims > 0 {
		writer.WriteHeader(http.StatusConflict)
		return
	}
	raceClaims++
	writer.WriteHeader(http.StatusOK)
}

func raceVulnerableClaim(writer http.ResponseWriter, _ *http.Request) {
	raceClaims++
	writer.WriteHeader(http.StatusOK)
}

func raceSafeOracle(writer http.ResponseWriter, _ *http.Request) {
	storeMutex.Lock()
	defer storeMutex.Unlock()
	if raceClaims > 1 {
		writer.WriteHeader(http.StatusConflict)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func raceVulnerableOracle(writer http.ResponseWriter, _ *http.Request) {
	if raceClaims > 1 {
		writer.WriteHeader(http.StatusConflict)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func registerServiceRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /controller/service/control", serviceWitness)
	mux.HandleFunc("POST /controller/service/safe", serviceSafe)
	mux.HandleFunc("POST /controller/service/vulnerable", serviceWitness)
}

func serviceSafe(writer http.ResponseWriter, _ *http.Request) { writer.WriteHeader(http.StatusOK) }

func serviceWitness(writer http.ResponseWriter, request *http.Request) {
	correlation := request.Header.Get("X-Rig-Protocol-Correlation")
	if correlation == "" {
		writer.WriteHeader(http.StatusOK)
		return
	}
	origin, capability := os.Getenv("RIG_BILLING_ORIGIN"), os.Getenv("RIG_BILLING_CAPABILITY")
	if origin == "" || capability == "" {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	witnessRequest, err := http.NewRequestWithContext(request.Context(), http.MethodPost, origin+"/charge", nil)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	witnessRequest.Header.Set("X-Rig-Witness-Capability", capability)
	witnessRequest.Header.Set("X-Rig-Protocol-Correlation", correlation)
	response, err := (&http.Client{Timeout: time.Second}).Do(witnessRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	_ = response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
