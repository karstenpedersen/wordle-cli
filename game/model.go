package game

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
