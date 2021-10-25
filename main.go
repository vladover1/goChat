package main

import (
	"fmt"
	"net/http"
)

func main() {

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/", fs)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/login.html")
	})

	fmt.Println("сервер ожидает")
	http.ListenAndServe(":8080", nil)
}
