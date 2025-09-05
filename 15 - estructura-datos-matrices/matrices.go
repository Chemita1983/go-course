package main

import "fmt"

func main() {
	var matriz [5]int // Nombre del array y longitud

	matriz[0] = 10
	matriz[1] = 20

	fmt.Println(matriz)

	var matriz2 = [5]int{1, 2, 3, 4, 5} // Otra manera de inicializar

	fmt.Println(matriz2)

	var matriz3 = [...]int{10, 20, 30, 40, 50} // Colocamos 3 puntos cuando no sabemos cuantos elementos tiene.

	fmt.Println(matriz3)

	// Recorrer la matrix
	for i := 0; i < len(matriz3); i++ {
		fmt.Println(matriz3[i])
	}

	for index, value := range matriz3 {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}

	var matrizBi = [2][2]int{{1, 2}, {3, 4}} // Matrices con mas de una dimension
	fmt.Println(matrizBi)
}
