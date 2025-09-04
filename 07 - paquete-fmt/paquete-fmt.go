/* Scan lee datos introducidos y recibe como parámetro la direccion de memoria
  de la variable donde va a guardar el valor, esta es con &{nombreVariable}
  podemos poner tantas variables como necesitemos, pero tienen que existir y
  haremos referencia a ellas

  	fmt.Print("Ingrese su nombre y apellido: ")
	fmt.Scanln(&firstName, &lastName)
*/

package main

import "fmt"

var (
	firstName string
	age       int
)

func main() {

	fmt.Print("Hola") // No hace salto de linea al final.
	fmt.Println("Hola Mundo")

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&firstName)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola me llamo %s y tengo %d años \n", firstName, age)

	greeting := fmt.Sprintf("Hola me llamo %s y tengo %d años", firstName, age) // Devuelve el resultado en string sin mostrarlo por pantalla
	fmt.Println(greeting)
	fmt.Printf("El timpo de dato de name es %T y el de age es %T\n", firstName, age) // %T muestra el tipo de dato
}
