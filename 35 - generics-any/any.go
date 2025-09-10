package main

import "fmt"

func main() {
	PrintList("Hola", 5, true, 4.5)

}

// Aqui usabamos antes interface{} en vez de any.
func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}
