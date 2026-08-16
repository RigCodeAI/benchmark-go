package main

import "net/http"

func main() {
	method := "GET"
	path := "/dynamic"
	http.HandleFunc(method+" "+path, func(http.ResponseWriter, *http.Request) {})
	_ = http.ListenAndServe(":8080", nil)
}
