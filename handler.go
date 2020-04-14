package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	users, err := retrieveA()
	if err != nil {
		errorMessage(w, r, "Cannot get Users")
	} else {
		// HTMLを生成して，ResponseWriterに書き出す
		generateHTML(w, users, "basis", "header", "index")
	}
}
