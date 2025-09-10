// Vamos a trabajar con tipos enteros en toda la funcion [T int] -> Restricción arbitraria
// En la siguiente función le indicamos que vamos a utilizar datos enteros o float
// ~ -> Con este simbolo creamos una restriccion de aproximación

package main

import "fmt"

type integer int // Con el simbolo ~ se hace una restriccion por aproximación y los valores integer son aceptados en la funcion

func Sum[T ~int | ~float64](nums ...T) T {
	var result T = 0
	for _, num := range nums {
		result += num
	}
	return result
}


func main() {

	var num integer = 100
	var num2 integer = 200
	fmt.Println(Sum(5, 6, 7, 8.5))
	fmt.Println(Sum(num, num2))

}
