package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type fixtureDatabase struct{}
type fixtureResult struct{}

func (fixtureDatabase) ExecContext(context.Context, string, ...any) (sql.Result, error) {
	return fixtureResult{}, nil
}
func (fixtureResult) LastInsertId() (int64, error) { return 0, nil }
func (fixtureResult) RowsAffected() (int64, error) { return 0, nil }

var database fixtureDatabase

func vulnerableSQL(writer http.ResponseWriter, request *http.Request) {
	value := request.URL.Query().Get("q")
	_, _ = database.ExecContext(request.Context(), "SELECT * FROM records WHERE name = '"+value+"'")
	writer.WriteHeader(http.StatusNoContent)
}

func safeSQL(writer http.ResponseWriter, request *http.Request) {
	value := request.URL.Query().Get("q")
	_, _ = database.ExecContext(request.Context(), "SELECT * FROM records WHERE name = ?", value)
	writer.WriteHeader(http.StatusNoContent)
}

func vulnerableCommand(writer http.ResponseWriter, request *http.Request) {
	value := request.URL.Query().Get("q")
	_ = exec.Command("sh", "-c", value).Run()
	writer.WriteHeader(http.StatusNoContent)
}

func safeCommand(writer http.ResponseWriter, request *http.Request) {
	_ = request.URL.Query().Get("q")
	_ = exec.Command("printf", "%s", "safe").Run()
	writer.WriteHeader(http.StatusNoContent)
}

func vulnerablePath(writer http.ResponseWriter, request *http.Request) {
	value := request.URL.Query().Get("q")
	_, _ = os.ReadFile(value)
	writer.WriteHeader(http.StatusNoContent)
}

func safePath(writer http.ResponseWriter, request *http.Request) {
	_ = request.URL.Query().Get("q")
	_, _ = os.ReadFile("go.mod")
	writer.WriteHeader(http.StatusNoContent)
}

func vulnerableXSS(writer http.ResponseWriter, request *http.Request) {
	value := request.URL.Query().Get("q")
	_, _ = fmt.Fprint(writer, value)
}

func safeXSS(writer http.ResponseWriter, request *http.Request) {
	_ = request.URL.Query().Get("q")
	_, _ = fmt.Fprint(writer, "safe")
}

func vulnerableSSRF(writer http.ResponseWriter, request *http.Request) {
	value := request.URL.Query().Get("q")
	response, _ := http.Get(value)
	if response != nil {
		_ = response.Body.Close()
	}
	writer.WriteHeader(http.StatusNoContent)
}

func safeSSRF(writer http.ResponseWriter, request *http.Request) {
	_ = request.URL.Query().Get("q")
	response, _ := http.Get("http://127.0.0.1:1/constant")
	if response != nil {
		_ = response.Body.Close()
	}
	writer.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("GET /vulnerable/sql", vulnerableSQL)
	http.HandleFunc("GET /safe/sql", safeSQL)
	http.HandleFunc("GET /vulnerable/command", vulnerableCommand)
	http.HandleFunc("GET /safe/command", safeCommand)
	http.HandleFunc("GET /vulnerable/path", vulnerablePath)
	http.HandleFunc("GET /safe/path", safePath)
	http.HandleFunc("GET /vulnerable/xss", vulnerableXSS)
	http.HandleFunc("GET /safe/xss", safeXSS)
	http.HandleFunc("GET /vulnerable/ssrf", vulnerableSSRF)
	http.HandleFunc("GET /safe/ssrf", safeSSRF)
	_ = http.ListenAndServe(":8080", nil)
}
