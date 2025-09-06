// En go un puntero es una variable que almacena la dirección de memoria de otra variable
// Se utilizan para referenciar y acceder a la variable original a través de su direccion de memoria

package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	email  string
}

// Metodos en los que se recibe una referencia a la clase Persona, accedemos a el mediante la instancia.
func (p *Persona) diHola() {
	fmt.Println("Hola mi nombre es", p.nombre)
}

func main() {

	var x int = 10
	//var p *int = &x // Esto es un puntero que recibe la dirección de memoria de la variable x

	fmt.Println(x)
	editar(&x)
	fmt.Println(x)

	p := Persona{"Chema", 42, "chema@gmail.com"}
	p.diHola()
}

// Modificamos el valor real de la variable, paso por referencia
func editar(x *int) {
	*x = 20
}

// Modificamos una copia, paso por valor
func editar2(x int) {
	x = 20
}
