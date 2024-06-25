package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/sid91520/mode1/pkg/models"
	"github.com/sid91520/mode1/pkg/utils"
)
var Newbook models.Book
func GetBook(w http.ResponseWriter,r *http.Request){
	newbooks:=models.GetBook()
	res,_:=json.Marshal(newbooks)
	w.Header().Set("Content-Type","pkglication/son")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	bookIDstr:=vars["bookID"]
	ID,err:=strconv.ParseInt(bookIDstr,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_:=models.GetBookById(ID)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res,_:=json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	bookIDstr:=vars["bookID"]
	ID,err:=strconv.ParseInt(bookIDstr,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	book:=models.DeleteBook(ID)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter,r *http.Request){
	var updatebook=models.Book{}
	utils.ParseBody(r,updatebook)
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookdetails,db:=models.GetBookById(Id)
	if updatebook.Name != ""{
		bookdetails.Name=updatebook.Name
	}
	if updatebook.Auhor != ""{
		bookdetails.Auhor=updatebook.Auhor
	}
	if updatebook.Publication !=""{
		bookdetails.Publication=updatebook.Publication
	}
	db.Save(&bookdetails)
	res,_:=json.Marshal(bookdetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}