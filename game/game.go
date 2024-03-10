package game

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
	"unicode"
)

func Start(word string) {
	if word == "" {
		word = "hello"
	}

	p := tea.NewProgram(initialModel(word), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
	}
}

func initialModel(word string) model {
	maxGuesses := len(word)
	return model{
		state:      GameState,
		word:       []rune(strings.ToUpper(word)),
		maxGuesses: maxGuesses,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case GameState:
		return m.GameUpdate(msg)
	default:
		return m.EndUpdate(msg)
	}
}

func (m model) GameUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.column == len(m.word) {
				m.guessed = string(m.word) == string(m.input)
				m.guesses = append(m.guesses, m.input)
				m.input = []rune{}
				m.line++
				m.column = 0

				if m.line == m.maxGuesses || m.guessed {
					m.state = EndState
					return m, tea.ClearScreen
				}
			}
		case "backspace":
			if m.column > 0 {
				m.input = m.input[:m.column-1]
				m.column--
			}
		default:
			r := unicode.ToUpper(msg.Runes[0])
			if unicode.IsLetter(r) && m.column < len(m.word) {
				m.input = append(m.input, r)
				m.column++
			}
		}
	}

	return m, nil
}

func (m model) EndUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case " ", "enter":
			return initialModel("word"), tea.ClearScreen
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.state {
	case GameState:
		return m.GameView()
	default:
		return m.EndView()
	}
}

func (m model) GameView() string {
	s := ""

	// Board
	for i := 0; i < m.maxGuesses; i++ {
		if i < m.line {
			for _, r := range m.guesses[i] {
				s += fmt.Sprintf("[%c]", r)
			}
		} else {
			for j := 0; j < len(m.word); j++ {
				if i == m.line && j < len(m.input) {
					s += fmt.Sprintf("[%c]", m.input[j])
				} else {
					s += "[ ]"
				}
			}
		}
		s += "\n"
	}

	return s
}

func (m model) EndView() string {
	s := fmt.Sprintf("The word was: %s\n\n", string(m.word))

	if m.guessed {
		try := "try"
		if m.line != 1 {
			try = "tries"
		}
		s += fmt.Sprintf("You guessed the word in %d %s.", m.line, try)
	} else {
		s += "You did not guess the word"
	}

	s += "\n\n"
	s += "Play: space\n"
	s += "Quit: ctrl+c"
	return s
}
