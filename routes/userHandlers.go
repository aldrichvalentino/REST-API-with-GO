package routes

import (
	"net/http"
	"fmt"
	"io/ioutil"
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
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	db, err := gorm.Open("mysql", "root:admin@tcp(db)/mydb?parseTime=true") 
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var rows model.User
	db.Find(&rows, id)
	if rows.ID == 0 {
		fmt.Println("Error: user not found")
		http.Error(w, "Error: user not found", http.StatusNotFound)
		return
	} 

	w.Header().Set("Content-Type", "application/json")
	users, _ := json.Marshal(rows)
	w.Write([]byte(users))
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var newUser model.User
	err := json.Unmarshal(body, &newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 
	
	db, err := gorm.Open("mysql", "root:admin@tcp(db)/mydb?parseTime=true") 
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create the user
	result := db.Create(&newUser)
	if result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	} 
	
	w.Header().Set("Content-Type", "application/json")
	responseJson := make(map[string]string)
	responseJson["message"] = "User created!"
	message, _ := json.Marshal(responseJson)
	w.Write([]byte(message))
	// TODO: make a token / session handling
}

func HandleEditUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	
	db, err := gorm.Open("mysql", "root:admin@tcp(db)/mydb?parseTime=true") 
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer db.Close()

	var user model.User
	db.Find(&user, id)
	if user.ID == 0 {
		fmt.Println("Error: user not found")
		http.Error(w, "user not found", http.StatusNotFound)
		return 
	} 
	
	body, _ := ioutil.ReadAll(r.Body)
	var userUpdates model.User
	err = json.Unmarshal(body, &userUpdates)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the user
	result := db.Model(&user).Updates(userUpdates)	
	if result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseJson := make(map[string]string)
	responseJson["message"] = "User updated!"
	message, _ := json.Marshal(responseJson)
	w.Write([]byte(message))
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	
	db, err := gorm.Open("mysql", "root:admin@tcp(db)/mydb?parseTime=true") 
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer db.Close()

	var user model.User
	db.Find(&user, id)
	if user.ID == 0 {
		fmt.Println("Error: user not found")
		http.Error(w, "user not found", http.StatusNotFound)
		return 
	} 

	// Delete the user
	result := db.Delete(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseJson := make(map[string]string)
	responseJson["message"] = "User deleted!"
	message, _ := json.Marshal(responseJson)
	w.Write([]byte(message))
}