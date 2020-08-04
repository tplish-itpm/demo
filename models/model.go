package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("sqlite3", "my.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Role{})

	db.CreateTable(&User{})
	db.CreateTable(&Role{})

	db.Create(&Role{RoleName: RoleUser})
	db.Create(&Role{RoleName: RoleAdmin})

	DB = db
}
