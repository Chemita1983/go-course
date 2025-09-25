package main

import (
	"apirest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mu"
)

const PORT = ":8080"

func main() {

	// Rutas
	mux := mux.NewRouter()

	// EndPoints
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// Inicia el servidor web en el puerto especificado.
	log.Fatal(http.ListenAndServe(PORT, mux))
}
