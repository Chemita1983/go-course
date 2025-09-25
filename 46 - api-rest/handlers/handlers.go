package handlers

import (
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("content-type", "application/json")
	users, _ := models.ListUsers()

	byteResult, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(byteResult))
}

func GetUser(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("content-type", "application/json")

	// Obtener id
	vars := mux.Vars(request)
	userdId, _ := strconv.Atoi(vars["id"])

	user, _ := models.GetUserById(userdId)

	byteResult, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(byteResult))
}

func CreateUser(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("content-type", "application/json")

	// Obtener el cuerpo de la request
	decoder := json.NewDecoder(request.Body)
	user := models.User{}

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		user.Save()
	}

	byteResult, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(byteResult))
}

func UpdateUser(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("content-type", "application/json")

	// Obtener id
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["id"])

	// Obtener el cuerpo de la request
	decoder := json.NewDecoder(request.Body)
	user := models.User{}

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		user.Id = int64(userId)
		user.Save()
	}

	byteResult, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(byteResult))
}

func DeleteUser(rw http.ResponseWriter, request *http.Request) {
	rw.Header().Set("content-type", "application/json")

	// Obtener id
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["id"])
	user, _ := models.GetUserById(userId)
	user.Delete()
}
