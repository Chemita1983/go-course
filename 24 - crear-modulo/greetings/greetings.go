/*
➜  24 - crear-modulo git:(main) ✗ cd greetings
➜  greetings git:(main) ✗ go mod init github.com/Chemita1983/greetings
go: creating new go.mod: module github.com/Chemita1983/greetings
*/

package greetings

import (
	"errors"
	"fmt"
)

// Devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty string name")
	}
	message := fmt.Sprintf("Hello %v", name)
	return message, nil
}
