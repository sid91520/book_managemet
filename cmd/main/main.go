// package main

// import(
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/sid91520/mode1/pkg/routes"
// )
// func main(){
// 	r:=mux.NewRouter()
// 	routes.RegisterBookStoreRouters(r)
// 	http.Handle("/",r)
// 	log.Fatal(http.ListenAndServe("localhost:9100",r))

// }
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sid91520/mode1/pkg/routes"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterBookStoreRouters(r)

	// Set up the server
	http.Handle("/", r)
	log.Println("Starting server on port 9100")

	// Start the server
	log.Fatal(http.ListenAndServe(":9100", r))
}
