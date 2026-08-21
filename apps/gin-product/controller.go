package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"sync"
	"time"
)

var storeMutex sync.Mutex
var raceClaims int
var workflowExecutions int

func registerControllerRoutes(router *gin.Engine) {
	router.POST("/controller/login", gin.WrapF(login))
	router.GET("/controller/csrf-token", gin.WrapF(csrfToken))
	router.GET("/controller/cwe-284/vulnerable", gin.WrapF(allow))
	router.GET("/controller/cwe-284/safe", gin.WrapF(allowIdentity("alice")))
	router.GET("/controller/cwe-287/vulnerable", gin.WrapF(allow))
	router.GET("/controller/cwe-287/safe", gin.WrapF(allowIdentity("alice")))
	router.GET("/controller/cwe-306/vulnerable", gin.WrapF(allow))
	router.GET("/controller/cwe-306/safe", gin.WrapF(allowIdentity("alice")))
	router.GET("/controller/cwe-639/vulnerable", gin.WrapF(allow))
	router.GET("/controller/cwe-639/safe", gin.WrapF(allowIdentity("alice")))
	router.GET("/controller/cwe-862/vulnerable", gin.WrapF(allow))
	router.GET("/controller/cwe-862/safe", gin.WrapF(allowIdentity("alice")))
	router.GET("/controller/cwe-863/vulnerable", gin.WrapF(allow))
	router.GET("/controller/cwe-863/safe", gin.WrapF(allowIdentity("admin")))
	router.POST("/controller/cwe-352/vulnerable", gin.WrapF(csrfVulnerable))
	router.POST("/controller/cwe-352/safe", gin.WrapF(csrfSafe))
	registerWorkflowRoutes(router)
	registerRaceRoutes(router)
	registerServiceRoutes(router)
}

func login(writer http.ResponseWriter, request *http.Request) {
	http.SetCookie(writer, &http.Cookie{Name: "sivere_session", Value: request.FormValue("username"), HttpOnly: true})
	writer.WriteHeader(http.StatusOK)
}

func csrfToken(writer http.ResponseWriter, _ *http.Request) {
	http.SetCookie(writer, &http.Cookie{Name: "sivere_csrf", Value: "sivere-token"})
	writer.WriteHeader(http.StatusOK)
}

func csrfVulnerable(writer http.ResponseWriter, _ *http.Request) { writer.WriteHeader(http.StatusOK) }

func csrfSafe(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-CSRF-Token") != "sivere-token" {
		writer.WriteHeader(http.StatusForbidden)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func allow(writer http.ResponseWriter, _ *http.Request) { writer.WriteHeader(http.StatusOK) }

func allowIdentity(identity string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("sivere_session")
		if err != nil || cookie.Value != identity {
			writer.WriteHeader(http.StatusForbidden)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func registerWorkflowRoutes(router *gin.Engine) {
	router.POST("/controller/workflow/control/reset", gin.WrapF(workflowReset))
	router.POST("/controller/workflow/control/execute", gin.WrapF(workflowLimited))
	router.POST("/controller/workflow/safe/reset", gin.WrapF(workflowReset))
	router.POST("/controller/workflow/safe/execute", gin.WrapF(workflowLimited))
	router.POST("/controller/workflow/vulnerable/reset", gin.WrapF(workflowReset))
	router.POST("/controller/workflow/vulnerable/execute", gin.WrapF(workflowUnlimited))
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

func registerRaceRoutes(router *gin.Engine) {
	router.POST("/controller/race/safe/reset", gin.WrapF(raceReset))
	router.POST("/controller/race/vulnerable/reset", gin.WrapF(raceReset))
	router.POST("/controller/race/safe/claim", gin.WrapF(raceSafeClaim))
	router.POST("/controller/race/vulnerable/claim", gin.WrapF(raceVulnerableClaim))
	router.GET("/controller/race/safe/oracle", gin.WrapF(raceSafeOracle))
	router.GET("/controller/race/vulnerable/oracle", gin.WrapF(raceVulnerableOracle))
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

func registerServiceRoutes(router *gin.Engine) {
	router.POST("/controller/service/control", gin.WrapF(serviceWitness))
	router.POST("/controller/service/safe", gin.WrapF(serviceSafe))
	router.POST("/controller/service/vulnerable", gin.WrapF(serviceWitness))
}

func serviceSafe(writer http.ResponseWriter, _ *http.Request) { writer.WriteHeader(http.StatusOK) }

func serviceWitness(writer http.ResponseWriter, request *http.Request) {
	correlation := request.Header.Get("X-Sivere-Protocol-Correlation")
	if correlation == "" {
		writer.WriteHeader(http.StatusOK)
		return
	}
	origin, capability := os.Getenv("SIVERE_BILLING_ORIGIN"), os.Getenv("SIVERE_BILLING_CAPABILITY")
	if origin == "" || capability == "" {
		writer.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	witnessRequest, err := http.NewRequestWithContext(request.Context(), http.MethodPost, origin+"/charge", nil)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	witnessRequest.Header.Set("X-Sivere-Witness-Capability", capability)
	witnessRequest.Header.Set("X-Sivere-Protocol-Correlation", correlation)
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
