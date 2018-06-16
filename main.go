package main

// native pkg
import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

// vendor pkg
import (
	"github.com/gorilla/mux"
)

// custom pkg
import (
	db "app/migrations"
	"app/routes"
)

type Page struct {
	Title string
	Body []byte
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, &Page{Title: title, Body: []byte("test")})
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body") // get with the name body
	fmt.Println(body)
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, &Page{Title: body, Body: []byte(body)})
}

func main() {
	// do migration
	db.Migrate()

	Router := mux.NewRouter()
	Router.HandleFunc("/", handleGet).Methods("GET")
	Router.HandleFunc("/", handlePost).Methods("POST")
	Router.HandleFunc("/user/", routes.HandleGetAllUser).Methods("GET")
	Router.HandleFunc("/user/{id:[0-9]+}", routes.HandleGetUserById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", Router))
}
