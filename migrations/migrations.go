package migrations

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"app/model"
)

func Migrate() {
	db, err := gorm.Open("mysql", "root:admin@tcp(db)/") 
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create database first
	db.Exec("CREATE DATABASE mydb")

	// use database
	db.Exec("USE mydb")

	// migrate
	db.AutoMigrate(&model.User{})

	fmt.Println("Database migraion success")
}