package config

import "example.com/hello/Desktop/go-workspace/pkg/mod/github.com/jinzhu/gorm@v1.9.12"

var (
	db *gorm.DB
)

func Connect() {
	prosgret_conname := ("host=db port=5432 user=postgres dbname=postgres password=example sslmode=disable")
	d, err := gorm.Open("postgres", prosgret_conname)
	//d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
