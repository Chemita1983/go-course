package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		integer16 int16 = 50
		integer32 int32 = 100

		floaat32 float32 = 10.2
		floaat64 float64 = 10.5
	)

	fmt.Println(int32(integer16) + integer32)
	fmt.Println(integer32 + int32(floaat64))
	fmt.Println(floaat32 + float32(integer32))

	sNumber := "5"
	valueConverted, _ := strconv.Atoi(sNumber) // paquete de go que convierte de string a integer entre otras cosas
	fmt.Println(valueConverted + 10)

	nInteger := 42
	sNumber = strconv.Itoa(nInteger) // convierte a string (recibe un int exclusivamente)
	fmt.Println(sNumber + " elementos")
	fmt.Println(sNumber + sNumber) // Se concatenan los strings
}
