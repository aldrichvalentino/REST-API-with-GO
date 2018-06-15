package migrations

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Migrate() {
	db, err := sql.Open("mysql", "root:admin@tcp(db)/") // no address means default localhost
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data
	err = db.Ping()
	if err != nil {
    panic(err.Error()) 
	}

	// migration
	db.Query("CREATE DATABASE mydb")
	db.Query("CREATE TABLE mydb.user(id int auto_increment, name varchar(255), age int, PRIMARY KEY(id))")
	// TODO: create seperate functions for creating db and tables

	fmt.Println("Database migraion success")
}