package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, request *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)
}

func GetUser(rw http.ResponseWriter, request *http.Request) {
	if user, err := getUserById(request); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

func CreateUser(rw http.ResponseWriter, request *http.Request) {

	user := models.User{}
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

func UpdateUser(rw http.ResponseWriter, request *http.Request) {
	var userId int64
	if old_user, err := getUserById(request); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = old_user.Id

		user := models.User{}
		decoder := json.NewDecoder(request.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, user, http.StatusOK)
		}
	}
}

func DeleteUser(rw http.ResponseWriter, request *http.Request) {
	if user, err := getUserById(request); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, "", http.StatusNoContent)
	}
}

func getUserById(request *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}
}
