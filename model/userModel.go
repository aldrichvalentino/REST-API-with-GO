package model

import (
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func GetUser() string {
	db, err := sql.Open("mysql", "root:admin@tcp(db)/mydb") 
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data
	err = db.Ping()
	if err != nil {
    panic(err.Error()) 
	}

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string
	var age int
	var users []User

	for rows.Next() {
		rows.Scan(&id, &name, &age)
		tempUser := User{
			Id: id,
			Name: name,
			Age: age,
		}
		users = append(users, tempUser)
	}
	result, _ := json.Marshal(users)
	return string(result)
}

func GetUserById(userId string) (string, error) {
	db, err := sql.Open("mysql", "root:admin@tcp(db)/mydb") 
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data
	err = db.Ping()
	if err != nil {
    panic(err.Error()) 
	}

	rows, err := db.Query("SELECT * FROM user WHERE id=" + userId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string
	var age int
	var user User

	for rows.Next() {
		rows.Scan(&id, &name, &age)
		user = User{
			Id: id,
			Name: name,
			Age: age,
		}
	}
	result, _ := json.Marshal(user)
	return string(result), nil // TODO: implement an error when not found
}

func createUser() {
	return
}

func updateUser() {
	return
}

func deleteUser() {
	return
}