package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sid91520/mode1/pkg/config"
)
var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Auhor      string `json:"author"`
	Publication string `json:"publication"`
}
// func init(){
// 	config.Connect()
// 	db=config.GetDB()
// 	db.AutoMigrate(&Book{})
// }
func init() {
    config.Connect()
    if config.GetDB() == nil {
        fmt.Println("Database connection is nil")
        panic("Database connection initialization failed")
    }
    fmt.Println("Database connection initialized successfully")
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBook() []Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB){
	var GetBookid Book
	db:=db.Where("ID=?",id).Find(&GetBookid)
	return &GetBookid,db
}

func DeleteBook(id int64) Book{
	var delebook Book
	db.Where("ID=?",id).Delete(&delebook)
	return delebook
}