package handlers

import (
	"apirest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, request *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, request *http.Request) {
	if user, err := getUserByRequest(request); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, request *http.Request) {
	
	user := models.User{}
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, request *http.Request) {

	var userId int64

	if user, err := getUserByRequest(request); err != nil {
		models.SendNotFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}
}

func DeleteUser(rw http.ResponseWriter, request *http.Request) {
	if user, err := getUserByRequest(request); err != nil {
		models.SendNotFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(request *http.Request) (models.User, error) {
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["id"])
	if user, err := models.GetUserById(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}

}
