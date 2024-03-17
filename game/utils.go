package game

import (
	"github.com/karstenpedersen/wordle-cli/utils"
	"math/rand"
	"os"
	"strings"
)

func readWordList(wordlistPath string) ([]string, error) {
  if wordlistPath == "" {
    return []string{}, nil
  }
	content, err := os.ReadFile(wordlistPath)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(content), "\n"), nil
}

func getRandomWord(wordlistPath string) (string, error) {
	lines, err := readWordList(wordlistPath)
	if err != nil {
		return "", err
	}
	randomLineIndex := rand.Intn(len(lines))
	word := lines[randomLineIndex]

	return word, nil
}

func isWordValid(word []rune, wordlist []string) bool {
	return utils.Contains(wordlist, strings.ToLower(string(word))) || len(wordlist) == 0
}
