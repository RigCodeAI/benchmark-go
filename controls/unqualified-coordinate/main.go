package main

import "net/http"

func main() {
	http.HandleFunc("GET /unsupported", func(http.ResponseWriter, *http.Request) {})
	_ = http.ListenAndServe(":8080", nil)
}
