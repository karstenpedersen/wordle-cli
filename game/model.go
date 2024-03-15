package game

import (
  "strings"
  "os"
)

type State int

const (
	GameState State = iota
	EndState
)

type model struct {
	state      State
	word       []rune
	guesses    [][]rune
	input      []rune
	maxGuesses int
	line       int
	column     int
	guessed    bool
}

func initialModel(word string) model {
	if word == "" {
		randomWord, err := getRandomWord()
		if err != nil {
			os.Exit(1)
		}
		word = randomWord
	}

	maxGuesses := len(word)
	return model{
		state:      GameState,
		word:       []rune(strings.ToUpper(word)),
		maxGuesses: maxGuesses,
	}
}

