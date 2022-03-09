package main

import (
	"log"
	"net/http"
	database "restapi/database"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "POST", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}

func main() {
	database.InitialMigration()
	initializeRouter()
}
