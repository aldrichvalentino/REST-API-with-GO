package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

// User schema
type User struct {
  gorm.Model
  Name         string `json:"name"`
  Age          sql.NullInt64 `gorm:"not null";json:"age"`
  Email        string  `gorm:"type:varchar(100);unique_index";json:"email"`
}
