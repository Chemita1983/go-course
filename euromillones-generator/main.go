package main

import (
	"fmt"
	"math/rand/v2"
)

var principales [5]int
var estrellas [2]int

func main() {
	EuromGenerator()
}

func EuromGenerator() {

	fmt.Println("Numeros generados para el €uroMillones")

	for i := range principales {
		for {
			numeroActual := rand.IntN(49) + 1

			if !existeNumero(numeroActual, principales[:]) {
				principales[i] = numeroActual
				break
			}
		}
	}

	for i := range estrellas {
		for {
			estrellaActual := rand.IntN(11) + 1

			if !existeNumero(estrellaActual, estrellas[:]) {
				estrellas[i] = estrellaActual
				break
			}
		}
	}

	for _, value := range principales {
		fmt.Print(" ", value)
	}
	fmt.Println("")

	for _, value := range estrellas {
		fmt.Print(" ", value)
	}
	fmt.Println("")

}

func existeNumero(numeroActual int, matriz []int) bool {
	for _, value := range matriz {
		if numeroActual == value {
			return true
		}
	}
	return false
}
