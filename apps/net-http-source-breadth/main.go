package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type sourcePayload struct {
	Value string `json:"value"`
}

func writeValue(writer http.ResponseWriter, value string) { _, _ = fmt.Fprint(writer, value) }

func querySource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, request.URL.Query().Get("value"))
}
func pathSource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, request.PathValue("value"))
}
func formSource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, request.FormValue("value"))
}
func jsonSource(writer http.ResponseWriter, request *http.Request) {
	var payload sourcePayload
	_ = json.NewDecoder(request.Body).Decode(&payload)
	writeValue(writer, payload.Value)
}
func headerSource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, request.Header.Get("X-Rig-Source"))
}
func cookieSource(writer http.ResponseWriter, request *http.Request) {
	cookie, _ := request.Cookie("rig_source")
	if cookie != nil {
		writeValue(writer, cookie.Value)
	}
}
func multipartSource(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseMultipartForm(1 << 20)
	writeValue(writer, request.FormValue("value"))
}
func bodySource(writer http.ResponseWriter, request *http.Request) {
	value, _ := io.ReadAll(request.Body)
	writeValue(writer, string(value))
}
func middlewareSource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, request.Header.Get("X-Rig-Middleware"))
}
func contextSource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, request.Header.Get("X-Rig-Context"))
}
func principalSource(writer http.ResponseWriter, request *http.Request) {
	writeValue(writer, strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer "))
}

func main() {
	http.HandleFunc("GET /sources/query", querySource)
	http.HandleFunc("GET /sources/path/{value}", pathSource)
	http.HandleFunc("POST /sources/form", formSource)
	http.HandleFunc("POST /sources/json", jsonSource)
	http.HandleFunc("GET /sources/header", headerSource)
	http.HandleFunc("GET /sources/cookie", cookieSource)
	http.HandleFunc("POST /sources/multipart", multipartSource)
	http.HandleFunc("POST /sources/body", bodySource)
	http.HandleFunc("GET /sources/middleware", middlewareSource)
	http.HandleFunc("GET /sources/context", contextSource)
	http.HandleFunc("GET /sources/principal", principalSource)
	_ = http.ListenAndServe(":8080", nil)
}
