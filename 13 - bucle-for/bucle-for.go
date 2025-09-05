package main

import "fmt"

func main() {

	/* Bucle infinito

	for {

	}

	*/

	// Bucle con condición

	var i int

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Mismo ejemplo anterior con el bucle típico

	for i = 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// continue: salta a la siguiente iteración
	// break: sale del bucle.

}
