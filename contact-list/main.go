/*
	Crearemos un gestor de contactos que guardará y mostrará contactos
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const FILE_NAME = "contacts.json"

type Contact struct {
	Name  string `json:"name"` // En el json de se mostrará en minusculas
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func main() {
	log.SetPrefix("main(): ")
	var contacts []Contact

	err := loadContactsFromFile(&contacts)
	if err != nil {
		log.Println("Error to load contacts: ", err)
	}

	for {
		fmt.Print(" ==== Contacts =====\n",
			"1. Add contact\n",
			"2. Show contacts\n",
			"3. Exit\n",
			"Chose an option: ",
		)

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			log.Println("Error to read option")
			return
		}

		switch option {
		case 1:
			reader := bufio.NewReader(os.Stdin)

			var newContact Contact
			fmt.Print("Name: ")
			newContact.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			newContact.Email, _ = reader.ReadString('\n')
			fmt.Print("Phone: ")
			newContact.Phone, _ = reader.ReadString('\n')

			contacts = append(contacts, newContact)
			if err = saveContactsToFile(contacts); err != nil {
				log.Println("Error write contacts in file: ", err)
			}
		case 2:
			showContacts(contacts)
		case 3:
			fmt.Println("Bye...")
			return
		default:
			fmt.Println("Incorrect option")
		}
	}
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts) // Serializa datos en un json
	if err != nil {
		return err
	}
	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts) // Le pasamos el puntero de nuestro slice donde queremos guardar los datos
	if err != nil {
		return err
	}

	return nil
}

func showContacts(contacts []Contact) {
	fmt.Println("==============================================")
	for index, contact := range contacts {
		fmt.Printf("%d. Nombre:  %s Email: %s Telefono: %s \n",
			index+1, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Println("==============================================")
}
