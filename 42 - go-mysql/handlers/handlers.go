package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/model"
	"log"
)

// Handler para listar los contactos
func GetConctacts(db *sql.DB) {
	// Consulta para obtener todos los contactos
	query := "SELECT * FROM contact"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	fmt.Println("\n LISTA DE CONTACTOS: ")
	fmt.Println("---------------------------------------------------------------------------------------------------")

	for rows.Next() {

		//Instancia de nuestro modelo contact
		contact := model.Contact{}

		var valueEmail sql.NullString // Para valores nulos
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		// Validamos que el valor no sea nulo
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "NONE"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("---------------------------------------------------------------------------------------------------")

	}
}

// Handler para obtener un contacto por id
func GetContactById(db *sql.DB, contactId int) {
	// Consulta para obtener un contacto por id
	query := "SELECT * FROM contact where id = ?"

	row := db.QueryRow(query, contactId)

	contact := model.Contact{}

	var valueEmail sql.NullString // Para valores nulos
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("⚠️ No se encontro ningún contacto con el id: %d", contactId)
		}
		log.Fatal(err)
	}

	// Validamos que el valor no sea nulo
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "NONE"
	}

	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)

}

// Handler para modificar un contacto
func UpdateContact(db *sql.DB, contact model.Contact) {
	// Consulta actuaizar un contacto por id
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Contacto modificado con exito")
}

// Handler para añadir un contacto
func AddContact(db *sql.DB, contact model.Contact) {
	// Consulta para añadir un contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?,?,?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Contacto registrado con exito")

}

// Handler para eliminar un contacto
func DeleteContact(db *sql.DB, contactId int) {
	// Consulta para eliminar un contacto por id
	query := "DELETE FROM contact WHERE id = ?"

	result, err := db.Exec(query, contactId)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		fmt.Printf("⚠️ No se encontró ningún contacto con el ID: %d\n", contactId)
		return
	}

	fmt.Println("✅ Contacto eliminado con éxito")
}
