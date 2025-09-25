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
	// Abrir una conexion de BBDD
	connection, err := sql.Open("mysql", getDBConnection())
	if err != nil {
		log.Fatal(err)
	}

	db = connection

	return nil
}

func Close() {
	db.Close()
}

func Ping() {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func ExistsTable(tableName string) bool {
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(query)
	if err != nil {
		log.Println("Error: ", err)
	}

	return rows.Next()
}

func CretateTable(schema string, tableName string) {
	if !ExistsTable(tableName) {
		_, err := Exec(schema)
		if err != nil {
			log.Println("Error: ", err)
		}
	}
}

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE TABLE `%s`", tableName)
	_, err := Exec(sql)
	if err != nil {
		log.Println("Error: ", err)
	}
}

// Polimorfismo de Exec
func Exec(query string, args ...any) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		log.Println("Error: ", err)
	}

	return result, err
}

// Polimorfismo de Query
func Query(query string, args ...any) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		log.Println("Error: ", err)
	}
	return rows, err
}

func getDBConnection() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	return dns

}
