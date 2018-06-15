package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"database/sql"
)

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gorilla/mux"
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
	http.Redirect(w, r, "/" + body, http.StatusFound)
}

func main() {
	// database
	db, err := sql.Open("mysql", "root:admin@tcp(db)/mydb") // no address means default localhost
	if err != nil {
		panic(err.Error()) // TODO: make propper error handling
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data
	err = db.Ping()
	if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("connection success!")
	}

	// testing queries
	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string

	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Printf("there is %d : %s \n", id, name)
	}

	http.HandleFunc("/", handleGet)
	http.HandleFunc("/save", handlePost)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
