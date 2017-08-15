package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lrw := newLoggingResponseWriter(w)
		t1 := time.Now()
		next.ServeHTTP(lrw, r)
		t2 := time.Now()
		statusCode := lrw.statusCode
		log.Printf("[%s] %q %d - %v\n", r.Method, r.URL.String(), statusCode, t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}
