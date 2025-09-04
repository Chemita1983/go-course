package main

import (
	"fmt"
	"math"
)

func main() {

	a := 10
	b := 3

	fmt.Println(a + b) // Suma
	fmt.Println(a - b) // Resta
	fmt.Println(a * b) // Multiplicacion
	fmt.Println(a / b) // Division
	fmt.Println(a % b) // Modulo

	// b++ incrementa en 1
	// b-- decrementa en 1

	c := 10
	c += 5 // c = c + 5
	fmt.Println(c)

	fmt.Println(math.Pi)

	p := math.Pow(2, 3) // Potencia  2e3
	fmt.Println(p)

	sqr := math.Sqrt(64) // Raiz cuadrada
	fmt.Println(sqr)
}
