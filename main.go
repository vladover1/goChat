package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello. Go to <a href='/login'>/login</a>"))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method: ", r.Method)

		hi := "nothing"

		if r.Method == "POST" {
			r.ParseForm()

			log.Println("Login: ", r.FormValue("login"))
			log.Println("Password: ", r.FormValue("pass"))

			hi = "this is POST request"
		}

		tmpl, _ := template.ParseFiles("./static/login.html")

		tmpl.Execute(w, hi)
	})

	log.Println("Listen and serve at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
