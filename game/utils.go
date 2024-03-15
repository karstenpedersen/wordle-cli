package game

import (
	"github.com/karstenpedersen/wordle-cli/utils"
	"math/rand"
	"os"
	"strings"
)

func readWordList() ([]string, error) {
	content, err := os.ReadFile("wordlist")
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(content), "\n"), nil
}

func getRandomWord() (string, error) {
	lines, err := readWordList()
	if err != nil {
		return "", err
	}
	randomLineIndex := rand.Intn(len(lines))
	word := lines[randomLineIndex]

	return word, nil
}

func isWordValid(word []rune) bool {
	words, err := readWordList()
	if err != nil {
		return false
	}
	return utils.Contains(words, strings.ToLower(string(word)))
}
