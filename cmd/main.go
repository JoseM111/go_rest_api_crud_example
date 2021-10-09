package main

import (
	. "log"
	. "net/http"
	
	"github.com/gorilla/mux"
	
	. "go_rest_api/pkg/handlers"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func main() {
	// ™ ____________
	router := mux.NewRouter()
	
	// GET request: GetAllBooks
	router.HandleFunc("/books", GetAllBooks).Methods(MethodGet)
	// GET request: GetBookByID
	router.HandleFunc("/books/{id}", GetBookByID).Methods(MethodGet)
	// POST request: AddBooks
	router.HandleFunc("/books", AddBooks).Methods(MethodPost)
	// PUT request: PutUpdateBook
	router.HandleFunc("/books/{id}", PutUpdateBook).Methods(MethodPut)
	// DELETE request: DeleteBookByID
	router.HandleFunc("/books/{id}", DeleteBookByID).Methods(MethodDelete)
	
	/* localhost:4000 */
	listeningPort(router)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func listeningPort(router *mux.Router) {
	Println("API is running on...http://localhost:4000")
	err := ListenAndServe(":4000", router)
	if err != nil {
		return
	}
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

