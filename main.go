package main

import (
	"context"
	"embed"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//go:embed static
	res   embed.FS
	pages = map[string]string{
		"/": "static/signin.html",
	}
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	log.Println("connect mongoDB")
	mongoURL := "mongodb://localhost:27017/admin"
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	collection = client.Database("register").Collection("signin")

}

type User struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
}

func getUserByLogin(login string) (User, error) {
	var u User
	if err := collection.FindOne(ctx, bson.M{
		"login": login,
	}).Decode(&u); err != nil {
		return u, err

	}
	return u, nil

}

func main() {

	// log.SetFlags(log.Ltime | log.Lshortfile)
	// //
	// // Parse command line
	// //
	// listenUrl := flag.String("url", "localhost:3000", "Host/port on which to run websocket listener")
	// mongoUrl := flag.String("mongo", "localhost", "URL of MongoDB server")
	// flag.Parse()
	// // Extract DB name from DB URL, if present
	// dbName := "tokenizer" // If no DB name specified, use "tokenizer"
	// switch _, auth, _, err := mgourl.ParseURL(*mongoUrl); true {
	// case err != nil:
	// 	log.Fatal("Could not parse MongoDB URL:", err)
	// case auth.Db != "":
	// 	dbName = auth.Db
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello. Go to <a href='/signin'>/signin</a>"))
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

	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Method: ", r.Method)
		tmpl, _ := template.ParseFiles("./static/signin.html")
		var tmplVar = map[string]string{}

		if r.Method == "POST" {
			r.ParseForm()

			log.Println("Login: ", r.FormValue("login"))
			log.Println("Password: ", r.FormValue("pass"))

			u, err := getUserByLogin(r.FormValue("login"))
			if err != nil {
				tmplVar["error"] = "user not found"
				tmpl.Execute(w, tmplVar)
				return
			}
			if u.Password != r.FormValue("pass") {
				tmplVar["error"] = "user not found"
				tmpl.Execute(w, tmplVar)
				return
			}

			http.Redirect(w, r, "/allok", 302)
		}
		tmpl.Execute(w, tmplVar)

	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./static/signup.html")

		tempVar := map[string]string{}
		if r.Method == "POST" {
			r.ParseForm()

			login := r.FormValue("login")
			pass := r.FormValue("pass")
			pass2 := r.FormValue("pass")

			if pass != pass2 {
				tempVar["error"] = "password not equal"
				tmpl.Execute(w, tempVar)
				return
			}

			existedUser, _ := getUserByLogin(login)
			if existedUser.Login == login {
				tempVar["error"] = "user login is not unique"
				tmpl.Execute(w, tempVar)
				return
			}

			_, err := collection.InsertOne(context.TODO(), User{
				Login:    login,
				Password: pass,
			})
			if err != nil {
				panic(err)
			}

			http.Redirect(w, r, "/", 302)
			return

		}
		tmpl.Execute(w, tempVar)
	})

	log.Println("Listen and serve at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
