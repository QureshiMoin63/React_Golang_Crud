package handlers

import (
	"log"
	"net/http"
	"restapi/docs"
	"restapi/models"

	"github.com/gorilla/handlers"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	//for JWT
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/login", models.Login).Methods("POST")
	r.HandleFunc("/home", models.Home).Methods("GET")
	r.HandleFunc("/refresh", models.Refresh).Methods("POST")
	r.HandleFunc("/logout", models.Logout).Methods("POST")
	// 	//For Books
	r.HandleFunc("/books/create", models.CreateBook).Methods("POST")
	r.HandleFunc("/books", models.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", models.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", models.UpdateBook).Methods("POST")
	r.HandleFunc("/books/{id}", models.DeleteBook).Methods("DELETE")
	//For USERS
	r.HandleFunc("/users", models.GetUser).Methods("GET")
	r.HandleFunc("/register", models.CreateUser).Methods("POST")
	// 	//For Authors
	r.HandleFunc("/authors", models.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", models.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", models.DeleteAuthor).Methods("DELETE")
	r.HandleFunc("/authors/{id}", models.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/create", models.CreateAuthor).Methods("POST")

	docs.SwaggerInfo.Host = "localhost:8080"
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "POST", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
	//log.Fatal(http.ListenAndServe(":8080",r))

	// }

}
