package main

import "fmt"

func main() {
	hello("Chema")
	fmt.Println(helloAsString("Chema"))
	sum, mul := calc(2, 3)
	sum2, mul2 := calc2(3, 6)

	fmt.Println("La suma es:", sum)
	fmt.Println("La multiplicacion es:", mul)

	fmt.Println("La suma2 es:", sum2)
	fmt.Println("La multiplicacion2 es:", mul2)

}

func hello(name string) {
	fmt.Println("Hello", name)
}

// Función que devuelve un unico valor
func helloAsString(name string) string {
	return "Hello " + name
}

// Devuelve multiples valores, los colocamos entre parentesis
func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

// Otra manera de hacerlo
func calc2(a, b int) (sum2, mul2 int) {
	sum2 = a + b
	mul2 = a * b
	return
}
