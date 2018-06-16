package routes

import (
	"net/http"
	"fmt"
	"strconv"
	"app/model"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func HandleGetAllUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:admin@tcp(db)/mydb?parseTime=true") 
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var rows []model.User
	db.Find(&rows)

	users, _ := json.Marshal(rows)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(users))
}

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	w.Header().Set("Content-Type", "application/json")

	db, err := gorm.Open("mysql", "root:admin@tcp(db)/mydb?parseTime=true") 
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var rows model.User
	db.Find(&rows, id)

	if rows.ID == 0 {
		fmt.Println("Error: user not found")
		error := make(map[string]string)
		error["message"] = "user not found!"
		message, _ := json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(message))
	} else {
		users, _ := json.Marshal(rows)
		w.Write([]byte(users))
	}
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blank"))
}

func HandleEditUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blank"))
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blank"))
}