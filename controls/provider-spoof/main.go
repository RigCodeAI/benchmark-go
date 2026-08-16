package main

import (
	"context"
	"net/http"
)

type spoofDatabase struct{}

func (spoofDatabase) ExecContext(context.Context, string, ...any) (any, error) { return nil, nil }

func main() {
	http.HandleFunc("GET /spoof", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = (spoofDatabase{}).ExecContext(request.Context(), request.URL.Query().Get("q"))
		writer.WriteHeader(http.StatusNoContent)
	})
	_ = http.ListenAndServe(":8080", nil)
}
