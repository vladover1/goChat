package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed static
	res   embed.FS
	pages = map[string]string{
		"/": "./static/login.html",
	}
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello. Go to <a href='/login'>/login</a>"))
		page, ok := pages[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		tpl, err := template.ParseFS(res, page)
		if err != nil {
			log.Printf("page %s not found in pages cache...", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"userAgent": r.UserAgent(),
		}
		if err := tpl.Execute(w, data); err != nil {
			return
		}
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
	http.FileServer(http.FS(res))

	log.Println("Listen and serve at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
