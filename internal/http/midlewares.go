package http

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipResponseWriter struct {
	http.ResponseWriter
	writer *gzip.Writer
}

func NewGzipResponseWriter(w http.ResponseWriter) *GzipResponseWriter {
	return &GzipResponseWriter{ResponseWriter: w}
}

func (w *GzipResponseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		gzWriter := NewGzipResponseWriter(w)
		gzWriter.writer = gz
		defer gzWriter.writer.Close()

		next.ServeHTTP(gzWriter, r)
	})
}
