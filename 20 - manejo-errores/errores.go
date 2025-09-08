package main

import (
	"fmt"
	"strconv"
)

func main() {

	// comoFuncionanErrores()

	resultado, error := divide(10, 0)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("Resultado: ", resultado)

}

func divide(a, b int) (int, error) {
	if b == 0 {
		//zeroDivisionError := errors.New("Error: Division by zero") // Creamos nuestros propios errores
		return 0, fmt.Errorf("Error: Division by zero: %d/%d\n", a, b) // Otra manera de crear errores para mostrar valores
	}

	return a / b, nil
}

func comoFuncionanErrores() {
	str := "123"
	num, err := strconv.Atoi(str)

	// err va a devolver dos datos si no hay error devuelve nil y si hay error devuelve un mensaje de
	// error
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Número: ", num)
}
