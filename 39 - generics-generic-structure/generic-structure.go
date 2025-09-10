package main

import "fmt"

// Un modelo de la base de datos (tabla)
type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {

	product1 := Product[uint]{1, "Producto 1", 50.0}
	product2 := Product[string]{"2", "Producto 2", 25.0}

	fmt.Println(product1)
	fmt.Println(product2)

}
