/*
 Funcion anonima -> funcion que no tiene nombre
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hola soy una función anónima")
	}()

	saludo := func(name string) {
		fmt.Println("Hola:", name)
	}

	//saludo("Chema")

	saludar("Chema", saludo)

	fmt.Println(apply(2, duplicar))
	fmt.Println(apply(2, triplicar))
}

func saludar(name string, f func(string)) {
	f(name)
}

func apply(n int, f func(int) int) int {
	return f(n)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}
