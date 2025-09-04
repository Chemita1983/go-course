/*
	 VALORES PREDETERMINADOS

		 Cuando declaramos una variable, esta tiene un valor por defecto.

		 int --> 0
		 uint --> 0
		 float32 --> 0
		 bool --> false
		 string --> ""
*/
package main

import "fmt"

func main() {

	var (
		defaultInt    int
		defaultUInt   uint
		defaulFloat32 float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println("Valores predeterminados: ")
	fmt.Println(defaultInt, defaultUInt, defaulFloat32, defaultBool, defaultString)
}
