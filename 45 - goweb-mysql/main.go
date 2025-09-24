package main

import (
	"gomysql/db"
	"log"
)

func main() {

	// Establecer la conexion con la BBDD
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
