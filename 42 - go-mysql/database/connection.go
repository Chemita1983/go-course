package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // El guion bajo indica que se va a usar de forma indirecta
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abrir una conexion de BBDD
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	// Comprobamos a conexion
	if err := db.Ping(); err != nil {
		return nil, err
	}

	//log.Println("Conexion a la base de datos sql exitosa")

	return db, nil

}
