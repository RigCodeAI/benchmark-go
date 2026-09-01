package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"testing"
)

var semanticCategoryIDs = []string{
	"cwe-113", "cwe-116", "cwe-1336", "cwe-200", "cwe-201", "cwe-22",
	"cwe-328", "cwe-330", "cwe-400", "cwe-501", "cwe-502", "cwe-532",
	"cwe-601", "cwe-611", "cwe-614", "cwe-643", "cwe-776", "cwe-78",
	"cwe-79", "cwe-89", "cwe-90", "cwe-918", "cwe-94", "cwe-943",
	"go-cgo-boundary", "go-goroutine-leak", "go-http-body-limit",
	"go-map-concurrent-access", "go-template-context-confusion",
}

func TestConcurrencyControlsHaveDeterministicOutcomes(t *testing.T) {
	handler := applicationHandler()
	for _, test := range []struct {
		name           string
		prefix         string
		workerStatuses []int
		oracleStatus   int
	}{
		{name: "safe", prefix: "/controller/race/safe", workerStatuses: []int{http.StatusOK, http.StatusConflict}, oracleStatus: http.StatusOK},
		{name: "vulnerable", prefix: "/controller/race/vulnerable", workerStatuses: []int{http.StatusOK, http.StatusOK}, oracleStatus: http.StatusConflict},
	} {
		t.Run(test.name, func(t *testing.T) {
			handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, test.prefix+"/reset", nil))
			start := make(chan struct{})
			statuses := make(chan int, 2)
			var workers sync.WaitGroup
			workers.Add(2)
			for range 2 {
				go func() {
					defer workers.Done()
					<-start
					response := httptest.NewRecorder()
					handler.ServeHTTP(response, httptest.NewRequest(http.MethodPost, test.prefix+"/claim", nil))
					statuses <- response.Code
				}()
			}
			close(start)
			workers.Wait()
			close(statuses)
			actual := []int{<-statuses, <-statuses}
			if actual[0] > actual[1] {
				actual[0], actual[1] = actual[1], actual[0]
			}
			if actual[0] != test.workerStatuses[0] || actual[1] != test.workerStatuses[1] {
				t.Fatalf("worker statuses = %v, want %v", actual, test.workerStatuses)
			}
			oracle := httptest.NewRecorder()
			handler.ServeHTTP(oracle, httptest.NewRequest(http.MethodGet, test.prefix+"/oracle", nil))
			if oracle.Code != test.oracleStatus {
				t.Fatalf("oracle status = %d, want %d", oracle.Code, test.oracleStatus)
			}
		})
	}
}

func TestAllProductControlRoutesAreExecutable(t *testing.T) {
	t.Cleanup(func() { _ = os.Remove("sivere-owned-go-deserialization-effect") })
	handler := applicationHandler()
	for _, category := range semanticCategoryIDs {
		for _, control := range []string{"vulnerable", "safe", "unknown"} {
			t.Run(category+"-"+control, func(t *testing.T) {
				value := "sivere"
				if category == "cwe-502" && control == "vulnerable" {
					value = `{"value":"sivere"}`
				}
				if category == "cwe-918" && control == "vulnerable" {
					value = "http://127.0.0.1:1/sivere"
				}
				request := httptest.NewRequest("GET", "/qualification/"+category+"/"+control+"?secret="+url.QueryEscape(value), nil)
				response := httptest.NewRecorder()
				handler.ServeHTTP(response, request)
				if response.Code == 404 {
					t.Fatalf("missing route for %s-%s", category, control)
				}
			})
		}
	}
}
