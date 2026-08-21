//go:generate sh -c "touch sivere-must-not-execute-generation"

package main

import "net/http"

func main() {
	http.HandleFunc("GET /generated", func(http.ResponseWriter, *http.Request) {})
	_ = http.ListenAndServe(":8080", nil)
}
