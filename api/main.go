package main

import (
	"fmt"
	//"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/items", GetItems)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func ConnectDB() {

}
