package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "password"
const DB_NAME = "bookstore"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var (
	db *gorm.DB
)

func Connect() {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
