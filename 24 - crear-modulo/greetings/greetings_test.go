package greetings

import (
	"regexp"
	"testing"
)

/*
		Las funciones de prueba reciben un parametro de tipo testing que se usa para reportar
	 	el resultado de la prueba
*/
func TestHelloName(t *testing.T) {
	name := "Chema"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Chema")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Chema") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
