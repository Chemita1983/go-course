// Proyecto para aprender las estructuras de control, bucles y demás.

/* Generarmos un numero aleatorio entre 0 y 100

le pediremos al usuario que intente averiguarlo durante 10 intentos

*/

package main

import (
	"fmt"
	"math/rand/v2"
)

const TRIES = 10

func main() {
	playGame()
}

func playGame() {

	var inputNumber int
	randomNumber := rand.IntN(100)

	for try := 1; try <= TRIES; try++ {
		fmt.Printf("Introduzca un numero: (intentos restantes %d): ", (TRIES - try + 1))
		fmt.Scanln(&inputNumber)

		if inputNumber == randomNumber {
			fmt.Println("Acertaste!!!!!")
			playGameAgain()
			return
		} else if inputNumber > randomNumber {
			fmt.Println("Incorrecto, El numero aleatorio es menor")
		} else {
			fmt.Println("Incorrecto. El numero aleatorio es mayor")
		}
	}

	fmt.Println("Lo sentimos, no acertaste, el numero era", randomNumber)
	playGameAgain()
}

func playGameAgain() {
	var option string

	fmt.Println("Desea jugar de nuevo: ")
	fmt.Scan(&option)

	switch option {
	case "s", "S":
		playGame()
	case "n", "N":
		fmt.Println("Gracias por jugar")
	default:
		playGameAgain()
	}
}
