/*
Closures -> una funcion anonima que se definida dentro de otra funcion y que ademas puede
modificar variables de la funcion externa y recuerda la definicion de una funcion ex
*/
package main

import "fmt"

func main() {
	nextInt := incrementer()
	fmt.Println(nextInt())

	// Si llamo varias veces recuerda el valor anterior
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4
	fmt.Println(nextInt()) // 5
	fmt.Println(nextInt()) // 6
	fmt.Println(nextInt()) // 7
}

func incrementer() func() int {
	i := 0 // Solo se hara esta sentencia cuando se llame a incrementer() en la linea 10
	return func() int {
		i++
		return i
	}
}
