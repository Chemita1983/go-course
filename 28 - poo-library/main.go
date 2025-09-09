/*
	Go no se basa fundamentalemte en POO pero si tiene algunas caracteristicas

	Pilares de la POO

		- Abstracción
		- Polimorfismo
		- Herencia
		- Encapsulamiento

		* Estructuras (Clases)
		* Métodos
		* Encapsulamiento
		* Composicion
		* Interfaces
*/

// Libreria
package main

import (
	"fmt"
	"poo/animal"
	"poo/book"
)

func main() {

	/*var book = book.Book{
		Title: "El Señor de los Anillos",
		Autor: "J.R.R Tolkien",
		Pages: 1200,
	}*/

	// Llamada al constructor que esta dentro de Book
	myBook := book.NewBook("El Señor de los Anillos", "J.R.R Tolkien", 1200)

	// Uso de set y get
	myBook.SetTitle("El Señor de los Anillos (Ed 2025)")
	fmt.Println(myBook.GetTitle())

	myTextBook := book.NewTextBook("Matematicas 1", "Varios autores", 300, "SM", "Primaria")

	// Probar polimorfismo
	book.Print(myBook)
	book.Print(myTextBook)

	fmt.Println("")

	// Guardamos estructuras de datos que cumplan con animal (Que implementen el metodo Sound())
	animals := []animal.Animal{
		&animal.Dog{Name: "Oscar"},
		&animal.Cat{Name: "Petunio"},
		&animal.Bird{Name: "Eustaquio"},
	}

	for _, animal := range animals {
		animal.Sound()
	}
}
