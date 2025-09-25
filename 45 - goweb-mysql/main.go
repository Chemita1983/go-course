package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"

	"log"
)

func main() {

	// Establecer la conexion con la BBDD
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	user := models.CreateUser("Marta C.H", "123456", "marta@email.com")
	user.Save()
	fmt.Println(models.ListUsers())
	db.Close()
}
