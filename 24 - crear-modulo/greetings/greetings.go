/*
➜  24 - crear-modulo git:(main) ✗ cd greetings
➜  greetings git:(main) ✗ go mod init github.com/Chemita1983/greetings
go: creating new go.mod: module github.com/Chemita1983/greetings
*/

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty string name")
	}
	message := fmt.Sprintf(randomMessage(), name)
	return message, nil
}

// Devuelve un saludo para varias personas especificadas
func Hellos(names []string) (map[string]string, error) {

	responseMessages := make(map[string]string)
	for _, name := range names {

		if name == "" {
			return responseMessages, errors.New("empty string name")
		}

		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		responseMessages[name] = message
	}
	return responseMessages, nil
}

// Devuelve un saludo aleaorio
func randomMessage() string {
	formats := []string{
		"Hello %v",
		"Hiiiiii %v!!!!!",
		"Nice to meet you %v",
	}

	return formats[rand.Intn(len(formats))]
}
