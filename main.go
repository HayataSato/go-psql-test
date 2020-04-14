package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//　DB接続
	psqlExec()

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)
	// REST likeなapi
	mux.HandleFunc("/api/", api)

	// starting up the server
	server := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        mux,
		ReadTimeout:    time.Duration(10 * int64(time.Second)),
		WriteTimeout:   time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("started at", server.Addr)
	server.ListenAndServe()
}
