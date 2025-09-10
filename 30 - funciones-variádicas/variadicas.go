/*
	Funciones variádicas son aquellas funciones que reciben n parametros
*/

package main

import "fmt"

type Estructura struct {
	Name string
}

func main() {

	fmt.Println("Result:", suma(1, 2, 3, 4, 5))
	est := Estructura{Name: "Chema"}
	imprimirDatos("Hola", 5, 4.5, true, nil, est)
}

// Enviamos como parámetros n argumentos
func suma(nums ...int) int {
	var result int = 0
	fmt.Printf("%T - %v\n", nums, nums)

	for _, num := range nums {
		result += num
	}
	return result
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}
