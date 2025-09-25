package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("content-type", "application/json")
	db.Connect()
	users := models.ListUsers()
	db.Close()

	byteResult, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(byteResult))
}

func GetUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "Listar un usuario")
}

func CreateUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}

func UpdateUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "Actualiza un usuario")
}

func DeleteUser(rw http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(rw, "Borra un usuario")
}
