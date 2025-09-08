/*
En go las palabras panic y recover se utilizan para manejar situaciones excepcionales o errores
graves que ocurren en la ejecución de un programa

panic -> generar una interrupcion normal del programa, cuando ocurre un error, el programa entra
en pánico.

recover -> se utiliza para capturar y manejar un pánico

Este uso se reserva para casos excepcionales o errores graves, no se recomienda su uso para un
flujo de datos normal.

Solo manejar situaciones inesperadas o tareas de limpieza al final de la ejecucion del programa
*/

package main

import "fmt"

func main() {
	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(100, 4)
}

func divide(a, b int) {
	defer func() { // Función anonima, con defer, recupera el pánico y no interrumpe la ejecución del programa
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validateZero(b)
	fmt.Println(a / b)
}

func validateZero(b int) {
	if b == 0 {
		panic("Error: Zero division") // Provocamos un pánico cuando ocurre un error
	}
}
