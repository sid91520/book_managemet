package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db *gorm.DB
)
// func Connect(){
// 	d,err:=gorm.Open("mysql", "sid:sid91520/simpleREST?charset=utf8&parseTime=True&loc=Local")
// 	if err !=nil{
// 		panic(err)
// 	} 
// 	db=d
// }

func Connect() {
d, err := gorm.Open("mysql", "root:sid91520@sharma@tcp(127.0.0.1:3306)/simpleREST?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    if err := d.DB().Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
    }

    db = d
    log.Println("Successfully connected to the database")
}
func GetDB() *gorm.DB{
	return db
}
