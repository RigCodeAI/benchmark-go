package main

import (
	"net/http/httptest"
	"net/url"
	"os"
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

func TestAllProductControlRoutesAreExecutable(t *testing.T) {
	t.Cleanup(func() { _ = os.Remove("rig-owned-go-deserialization-effect") })
	handler := applicationHandler()
	for _, category := range semanticCategoryIDs {
		for _, control := range []string{"vulnerable", "safe", "unknown"} {
			t.Run(category+"-"+control, func(t *testing.T) {
				value := "rig"
				if category == "cwe-502" && control == "vulnerable" {
					value = `{"value":"rig"}`
				}
				if category == "cwe-918" && control == "vulnerable" {
					value = "http://127.0.0.1:1/rig"
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
