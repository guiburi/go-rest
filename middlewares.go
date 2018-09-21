package main

import (
	"net/http"
	"log"
	"time"
)

func Logger(l *log.Logger) func(http.Handler) http.Handler {
	f := func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start:= time.Now()
			h.ServeHTTP(w, r)
			l.Printf("time taken: %v", time.Since(start))
		}
		return http.HandlerFunc(fn)
	}
	return f
}
