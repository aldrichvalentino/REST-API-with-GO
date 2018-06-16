package model

import (
	"github.com/jinzhu/gorm"
)

// User schema
type User struct {
  gorm.Model
  Name         string `json:"name"`
  Age          int `gorm:"not null";json:"age"`
  Email        string  `gorm:"type:varchar(100);unique_index";json:"email"`
}
