package server

import (
	"encoding/json"
	"io"
	"net/http"
)

func SendJSON(writer http.ResponseWriter, value interface{}, code int) {
	writer.Header().Add("Content-Type", "application/json; charset=utf-8")

	b, err := json.Marshal(value)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(writer, `{"error": "Internal server error"}`)

		return
	}

	writer.WriteHeader(code)
	_, _ = writer.Write(b)
}
