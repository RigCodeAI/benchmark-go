package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type sourceDocument struct {
	Value string `json:"value"`
}

func observeSource(writer http.ResponseWriter, value string) {
	encoded := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", "\"", "&quot;", "'", "&#39;").Replace(value)
	_, _ = fmt.Fprint(writer, encoded)
}

func registerSourceRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /sources/query", func(w http.ResponseWriter, r *http.Request) { observeSource(w, r.URL.Query().Get("value")) })
	mux.HandleFunc("GET /sources/path/{value}", func(w http.ResponseWriter, r *http.Request) { observeSource(w, r.PathValue("value")) })
	mux.HandleFunc("POST /sources/form", func(w http.ResponseWriter, r *http.Request) { observeSource(w, r.FormValue("value")) })
	mux.HandleFunc("POST /sources/json", func(w http.ResponseWriter, r *http.Request) {
		var value sourceDocument
		_ = json.NewDecoder(r.Body).Decode(&value)
		observeSource(w, value.Value)
	})
	mux.HandleFunc("GET /sources/header", func(w http.ResponseWriter, r *http.Request) { observeSource(w, r.Header.Get("X-Sivere-Source")) })
	mux.HandleFunc("GET /sources/cookie", func(w http.ResponseWriter, r *http.Request) {
		value, _ := r.Cookie("sivere_source")
		if value != nil {
			observeSource(w, value.Value)
		}
	})
	mux.HandleFunc("POST /sources/multipart", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseMultipartForm(1 << 20)
		observeSource(w, r.FormValue("value"))
	})
	mux.HandleFunc("POST /sources/body", func(w http.ResponseWriter, r *http.Request) {
		value, _ := io.ReadAll(http.MaxBytesReader(w, r.Body, 64<<10))
		observeSource(w, string(value))
	})
	mux.HandleFunc("GET /sources/middleware", func(w http.ResponseWriter, r *http.Request) { observeSource(w, r.Header.Get("X-Sivere-Middleware")) })
	mux.HandleFunc("GET /sources/context", func(w http.ResponseWriter, r *http.Request) { observeSource(w, r.Header.Get("X-Sivere-Context")) })
	mux.HandleFunc("GET /sources/principal", func(w http.ResponseWriter, r *http.Request) {
		observeSource(w, strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
	})
}
