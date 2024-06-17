package router

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

func gzipHandlerMiddleware(pass http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if strings.Contains(request.Header.Get("Content-Encoding"), "gzip") {
			reader, err := gzip.NewReader(request.Body)
			if err != nil {
				_, _ = io.WriteString(writer, `{"error": "Internal server error"}`)
				return
			}
			defer reader.Close()

			request.Header.Del("Content-Encoding")
			request.Header.Del("Content-Length")

			request.Body = reader
		}

		var w http.ResponseWriter = writer

		if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			gzipWriter := gzip.NewWriter(writer)
			defer gzipWriter.Close()

			w = &customResponseWriter{
				ResponseWriter: writer,
				writer:         gzipWriter,
			}

			writer.Header().Set("Content-Encoding", "gzip")
			writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			writer.Header().Set("Vary", "Accept-Encoding")
		}

		pass.ServeHTTP(w, request)
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	writer *gzip.Writer
}

func (c *customResponseWriter) Write(p []byte) (n int, err error) {
	return c.writer.Write(p)
}
