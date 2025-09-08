/*
Defer se utiliza para posponer la ejecucion de una funcion hasta que la funcion que la contiene
haya finalizado.

# Si se ponen defer en todas varias funciones, las funciones defer se agregan a una pila de ejecucion

(Ultimo en llegar, primero en ejecutarse, primero en llegar ultimo en ejecutarse)
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	_, err = file.Write([]byte("Hola Chema"))
	if err != nil {
		fmt.Println("Error escritura: ", err)
		return
	}

	defer file.Close() // Cerramos el flujo, con defer se va a ejecutar al final 

}

func usoDefer() {
	defer fmt.Println(3) // Se va a ejecutar al final de la funcion main
	fmt.Println(2)
	fmt.Println(1)
}
