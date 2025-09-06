// Struct un tipo de dato personalizado definido por tipos y valores especificos
package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	email  string
}

func main() {

	// Declaracion
	var p Persona
	p.nombre = "Chema"
	p.edad = 42
	p.email = "email@gmail.com"
	fmt.Println(p)

	p2 := Persona{"Juan", 50, "emailjuan@gmail.com"}
	fmt.Println(p2)
}
