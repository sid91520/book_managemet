// package routes

// import (
// 	"github.com/gorilla/mux"
// 	"github.com/sid91520/mode1/pkg/controllers"
// )
// var RegisterBookStoreRouters=func (router *mux.Router)  {
// 	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
// 	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
// 	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
// 	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
// 	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("GET")
// }

package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterBookStoreRouters(router *mux.Router) {
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all books"))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a single book"))
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new book"))
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a book"))
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a book"))
}
