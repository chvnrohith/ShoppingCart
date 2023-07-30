package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//prosgret_conname := ("host=db port=5432 user=postgres dbname=postgres password=example sslmode=disable")
	d, err := gorm.Open("mysql", "root:root@/project?charset=utf8&parseTime=True&loc=Local")
	//d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
