package main

import (
	db_mongo "api/db/mongo"
	redis_db "api/db/redis"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := db_mongo.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect db")
		return
	}
	err = redis_db.NewClient()
	if err != nil {
		log.Fatal("Failed to connect redis")
		return
	}
	log.Println("Connected to redis...")
	log.Println("Connected to mongo...")
	mux := mux.NewRouter().StrictSlash(false)
	log.Println("listening on Port: 8080..., localhost:http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}