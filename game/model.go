package game

import (
	"math/rand"
	"strings"
)

type State int

const (
	GameState State = iota
	EndState
)

type model struct {
	state      State
	word       []rune
	wordlist   []string
	guesses    [][]rune
	input      []rune
	maxGuesses int
	line       int
	column     int
	guessed    bool
}

func processWord(word string, wordlist []string) string {
	word = strings.TrimSpace(word)
	if word == "" && len(wordlist) > 0 {
		randomLineIndex := rand.Intn(len(wordlist))
		word = wordlist[randomLineIndex]
	}
	return word
}

func initialModel(word string, wordlist []string) model {
	maxGuesses := len(word)
	return model{
		state:      GameState,
		word:       []rune(strings.ToUpper(word)),
		maxGuesses: maxGuesses,
		wordlist:   wordlist,
	}
}
