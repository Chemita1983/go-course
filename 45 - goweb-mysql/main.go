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
	fmt.Println(models.ListUsers())
	db.Close()
}
d