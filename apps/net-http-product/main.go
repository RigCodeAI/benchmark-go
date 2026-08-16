package main

import "net/http"

func applicationHandler() http.Handler {
	mux := http.NewServeMux()
	registerSourceRoutes(mux)
	registerExplicitSemanticRoutes(mux)
	registerControllerRoutes(mux)
	return mux
}

func main() {
	_ = http.ListenAndServe(":8080", applicationHandler())
}
