package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type LogResponseWriter struct {
	http.ResponseWriter
	Request    string
	StatusCode int
	Buf        bytes.Buffer
}

func NewLogResponseWriter(r *http.Request, w http.ResponseWriter) *LogResponseWriter {
	body, _ := ioutil.ReadAll(r.Body)
	return &LogResponseWriter{
		ResponseWriter: w,
		Request:        string(body),
	}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.Buf.Write(body)
	return w.ResponseWriter.Write(body)
}
