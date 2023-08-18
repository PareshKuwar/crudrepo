// main/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// Change this to your actual project path
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id:[0-9]+}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/items/{id:[0-9]+}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id:[0-9]+}", handlers.DeleteItem).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
