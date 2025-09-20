package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/model"
	"log"
	"os"
	"strings"
)

const UPDATE_CONTACT_OPTION = 4

func main() {

	// Establecer la conexion con la BBDD
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("----- CONTACTOS -----")
		fmt.Println("1 - Mostrar todos contactos ")
		fmt.Println("2 - Mostrar un contacto por id")
		fmt.Println("3 - Agregar un contacto")
		fmt.Println("4 - Modificar un contacto")
		fmt.Println("5 - Eliminar un contacto")
		fmt.Println("6 - Salir")
		fmt.Print("Introduzca una opcion: ")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.GetConctacts(db)

		case 2:
			var id int
			fmt.Print("Introduzca el id del contacto: ")
			fmt.Scanln(&id)
			handlers.GetContactById(db, id)

		case 3:
			newContact := inputContactDetails(option)
			handlers.AddContact(db, newContact)
			handlers.GetConctacts(db)

		case 4:
			updatedContact := inputContactDetails(option)
			handlers.UpdateContact(db, updatedContact)
			handlers.GetConctacts(db)

		case 5:
			var id int
			fmt.Print("Introduzca el id del contacto para eliminar: ")
			fmt.Scanln(&id)
			handlers.DeleteContact(db, id)
		case 6:
			fmt.Println("Saliendo de contactos....")
			return

		default:
			fmt.Println("Opción no implementada")
		}
	}
}

func inputContactDetails(option int) model.Contact {

	reader := bufio.NewReader(os.Stdin)

	var newContact model.Contact

	if option == UPDATE_CONTACT_OPTION {
		var idContact int
		fmt.Print("Introduzca el id del contacto: ")
		fmt.Scanln(&idContact)
		newContact.Id = idContact
	}

	fmt.Print("Introduzca el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	newContact.Name = strings.TrimSpace(name)

	fmt.Print("Introduzca el email del contacto: ")
	email, _ := reader.ReadString('\n')
	newContact.Email = strings.TrimSpace(email)

	fmt.Print("Introduzca el telefono del contacto: ")
	phone, _ := reader.ReadString('\n')
	newContact.Phone = strings.TrimSpace(phone)

	return newContact
}
