/* Numeros
	var numero1 int8 //  8 bits (-128 a 127)
	var numero2 int // Depende de nuestro S.O
 	var numero3 uint // Solo numeros enteros positivos, depende tambien del S.O

 	Numeros decimales
 	var float1 float32 // Numeros muy grandes

	Booleanos
	var isOk bool = true

	Byte almacena numeros sin signo desde 0 hasta 255, se utiliza comunmente
	para representar datos de caracter ASCII y otros datos de tipo Byte
	var a byte = 'a' -> 97 el valor de a en formato ASCII

	Otro ejemplo es:

		s := "hola"
		fmt.Println(s[0])

	Esto imprimira el primer valor en decimal de la cadena "hola",
	en este caso la h (104)

	Para imprimir el valor unicode utilizamos rune que es un alias de int32

		var emoticono rune = '😎'
		fmt.Println(emoticono)
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Valor mínimo y máximo de Int64")
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	fmt.Println("Valor máximo de Float32")
	fmt.Println(math.MaxFloat32)

	fmt.Println("Valor máximo de Float64")
	fmt.Println(math.MaxFloat64)

	var isOk bool = true
	fmt.Println(isOk)

	fullName := "Chema J.V \t(alias \"Fean\")\n" // \t -> tabulador \ -> escapa caracteres \n -> salto de linea
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	var emoticono rune = '😎'
	fmt.Println(emoticono)
}
