package main

import (
	"net/http"
	"time"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hi there"))
	})
	server := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  3 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	func() {
		server.ListenAndServe()
	}()
}
