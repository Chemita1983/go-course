/*
 https://pkg.go.dev/golang.org/x/exp/constraints -> restricciones predefinidas
*/

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	var num int = 100
	var num2 int = 200
	fmt.Println(Sum(5, 6, 7, 8.5))
	fmt.Println(Sum(num, num2))
}

// Podemos crear nuestro tipo personalizado de restrición
//type Numbers interface {
//	~int | ~uint | ~float32 | ~float64
//}

// Aplicamos nuestro tipo personalizado de restricción
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var result T = 0
	for _, num := range nums {
		result += num
	}
	return result
}
