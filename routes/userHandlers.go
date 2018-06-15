package routes

import (
	"net/http"
	"app/model"
	"github.com/gorilla/mux"
)

// TODO: implement a not found response

func HandleGetAllUser(w http.ResponseWriter, r *http.Request) {
	users := model.GetUser()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(users))
}

func HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	users, _ := model.GetUserById(vars["id"])
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(users))
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