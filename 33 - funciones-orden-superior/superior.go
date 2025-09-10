/*
	Las funciones de orden superior -> Funcion que recibe como argumento, una funcion y devuelve
	otra función, es una programación mas modular y personalizada
*/

package main

import "fmt"

func main() {
	r := double(addOne, 3)
	fmt.Println("Result:", r)
}

func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}
