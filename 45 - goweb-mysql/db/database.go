package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() error {

	err := godotenv.Load()
	if err != nil {
		return err
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abrir una conexion de BBDD
	connection, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	db = connection
	log.Println("Conexion a la base de datos sql exitosa")

	return nil
}

func Close() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
