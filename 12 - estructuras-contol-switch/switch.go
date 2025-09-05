package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	switch os := runtime.GOOS; os { // Podemos declarar la variable dentro del switch
	case "darwin":
		fmt.Println("Estas en Mac")
	case "windows":
		fmt.Println("Estas en Windows")
	case "linux":
		fmt.Println("Estas en Linux")
	default:
		fmt.Println("S.O Desconocido")
	}

	// También podemos colocar condiciones booleanas dentro de un switch (con el ejemplo de la parte 11)

	switch tiempo := time.Now().Hour(); { // Si hay condiciones no se pone al final el valor a evaluar
	case tiempo < 13:
		fmt.Println("Es por la mañana")
	case tiempo < 20:
		fmt.Println("Es por la tarde")
	default:
		fmt.Println("Es por la noche")
	}
}
