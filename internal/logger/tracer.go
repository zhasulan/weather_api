package logger

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{w, 200, bytes.Buffer{}}
}

type responseRecorder struct {
	http.ResponseWriter
	Status   int
	BodyBuff bytes.Buffer
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.BodyBuff.Write(b)
	return rec.ResponseWriter.Write(b)
}

func (rec *responseRecorder) GUnZipBody() string {
	reader, err := gzip.NewReader(bytes.NewReader(rec.BodyBuff.Bytes()))
	if err != nil {
		return ""
	}

	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return ""
	}

	return string(body)
}

func (rec *responseRecorder) Body() string {
	return string(rec.BodyBuff.Bytes())
}
