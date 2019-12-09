package main

import (
	"context"
	"encoding/json"
	"time"

	//"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Site : An instance of site contains the name, URL and ClickCount for that site
type Site struct {
	Name       string
	URL        string
	ClickCount int
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/item", PostItem).Methods("POST")
	corsObj := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(corsObj)(router)))
}

// GetItems : Return All the Sites  in the Database
func GetItems(w http.ResponseWriter, r *http.Request) {
	var sites []Site
	collection := ConnectDB("sites")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Error: %s", err)
	}
	for cur.Next(context.Background()) {
		var site Site
		cur.Decode(&site)
		sites = append(sites, site)
		json.NewEncoder(w).Encode(sites)
	}
}

// PostItem : Function to add a Site to the Database
func PostItem(w http.ResponseWriter, r *http.Request) {
	var site Site
	_ = json.NewDecoder(r.Body).Decode(&site)
	w.Header().Set("content-type", "application/json")
	collection := ConnectDB("sites")
	result, _ := collection.InsertOne(context.Background(), site)
	json.NewEncoder(w).Encode(result)

}

// ConnectDB : Function to connect to the database, takes the name of the collection
func ConnectDB(collectionName string) *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Printf("Error %s", err)
		return nil
	}
	collection := client.Database(collectionName).Collection("sites")
	return collection
}
