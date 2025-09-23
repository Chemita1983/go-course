package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Piedra. Vence a tijeras. (tijeras + 1) % 3 = 0
	PAPER    = 1 // Papel. Vence a piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 // Tijeras. Vence a papel. (papel + 1) % 3 = 2
)

// Estructura para dar resultado de cada ronda
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Ganaste! 🎉",
	"Buen trabajo!",
	"¡Eres un campeón!",
}

var loseMessages = []string{
	"¡Perdiste! 😞",
	"No te rindas, sigue intentando.",
	"¡Casi lo logras!",
}

var drawMessages = []string{
	"¡Empate! 🤝",
	"¡Estás igualado!",
	"Nadie gana, nadie pierde.",
}

var ComputerScore, PlayerScore int

func PlayRound(playerChoice int) Round {
	computerValue := rand.Intn(3) // Genera un número aleatorio entre 0 y 2
	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La cpu eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La cpu eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La cpu eligió TIJERAS"
	}

	// Generar un número aleatorio para seleccionar el mensaje
	messageInt := rand.Intn(3)

	var message string

	if playerChoice == computerValue {
		roundResult = "Empate"
		message = drawMessages[messageInt]
	} else if playerChoice == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "Ganaste"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "Perdiste"
		message = loseMessages[messageInt]
	}

	return Round{
			Message: message, 
			ComputerChoice: computerChoice, 
			RoundResult: roundResult, 
			ComputerChoiceInt: computerChoiceInt, 
			ComputerScore: strconv.Itoa(ComputerScore), 
			PlayerScore: strconv.Itoa(PlayerScore),
		}
}
