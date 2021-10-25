package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8080", nil)
}
