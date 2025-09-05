package main

import "fmt"

func main() {

	// Comparación de números
	fmt.Println(1 == 1) // true
	fmt.Println(1 != 2) // true
	fmt.Println(2 < 3)  // true
	fmt.Println(3 > 4)  // false
	fmt.Println(4 <= 4) // true
	fmt.Println(5 >= 6) // false

	// Comparación de cadenas
	fmt.Println("hola" == "hola")  // true
	fmt.Println("hola" != "adios") // true
	fmt.Println("abc" < "def")     // true
	fmt.Println("ghi" > "fgh")     // true
	fmt.Println("hij" <= "hij")    // true
	fmt.Println("klm" >= "klmno")  // false

	// Comparación de booleanos
	fmt.Println(true == true)           // true
	fmt.Println(false != true)          // true
	fmt.Println(true && false == false) // true
	fmt.Println(true || false == true)  // true
}
