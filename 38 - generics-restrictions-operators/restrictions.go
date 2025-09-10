package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {

	videogames := []string{"Elden Ring", "Silksong", " Shinobi", "Street Fighter"}
	numbers := []int{1, 2, 3, 10, 12, 42, 46}

	if exist, videogame := Includes(videogames, "Silksong"); exist {
		fmt.Printf("%s exist\n", videogame)
	} else {
		fmt.Printf("%s not exist\n", videogame)
	}

	if exist, number := Includes(numbers, 30); exist {
		fmt.Printf("%d exist\n", number)
	} else {
		fmt.Printf("%d not exist\n", number)
	}

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
}

// Con comparable utilizamos dentro de la funcion operadores como = o !=
func Includes[T comparable](list []T, value T) (bool, T) {
	for _, acutalListValue := range list {
		if value == acutalListValue {
			return true, value
		}
	}
	return false, value

}

func Filter[O constraints.Ordered](list []O, callback func(O) bool) []O {
	result := make([]O, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}

	}
	return result
}
