package dependenciesinjection

import (
	"fmt"
	"io"
	"net/http"
)

// logMessage can write to any io.Writer without modification
func logMessage(w io.Writer, level, msg string) {
	fmt.Fprintf(w, "[%s] %s", level, msg)
}

func LogMessageHandler(w http.ResponseWriter, req *http.Request) {
	msg := req.URL.Query().Get("message")
	level := req.URL.Query().Get("level")
	if msg == "" {
		http.Error(w, "Missing message parameter", http.StatusBadRequest)
		return
	}
	if level == "" {
		level = "INFO"
	}
	logMessage(w, level, msg)
}
