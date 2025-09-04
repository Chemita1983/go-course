/*
	Una constante es una variable con un valor que no cambia.

	Se recomienda crearlas a nivel de paquete.

	Si no inicializamos la constante el programa nos deja ejecutar al
	contrario que con las variables.

	Podemos declarar varias costantes con:

	var (
		X = 100
		Y = 0b1010 // Binario
		Z = 0o12 // Octal
		W = 0xFF // Hexadecimal
	)


*/

package main

import "fmt"

const PI float32 = 3.1416

const (
	X = 100
	Y = 0b1010 // Binario
	Z = 0o12   // Octal
	W = 0xFF   // Hexadecimal
)

const (
	LUNES = iota + 1 // iota es un constante que representa 0, los demas valores se van incrementando en 1
	MARTES
	MIERCOLES
	JUEVES
	VIERNES
	SABADO
	DOMINGO
)

func main() {

	fmt.Println(X, Y, Z, W)
	fmt.Println(LUNES)

}
