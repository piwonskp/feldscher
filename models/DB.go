package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const ConnectionString = "host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"

var DB = Connect()

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		panic(err)
	}
	return db
}
