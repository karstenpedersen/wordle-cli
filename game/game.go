package game

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/karstenpedersen/wordle-cli/utils"
	"golang.org/x/term"
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

func ScreenView(s string) string {
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	return screenStyle.Width(width).Height(height).Render(s)
}

func (m model) GameView() string {
	s := BoardView(m)
	return ScreenView(s)
}

func BoardView(m model) string {
	s := ""
	for i := 0; i < m.maxGuesses; i++ {
		style := lineStyle
		if i == m.line {
			style = cursorLineStyle
		}
		s += style.Render(Line(m, i))
    s += "\n"
	}

	return boardStyle.Render(s)
}

func Line(m model, index int) string {
	s := ""
	if index < m.line {
		for i, r := range m.guesses[index] {
			style := incorrectTileStyle
			if r == m.word[i] {
				style = correctTileStyle
			} else if utils.Contains(m.word, r) {
				style = presentTileStyle
			}
			s += style.Render(string(r))
		}
	} else if index == m.line {
		for i := 0; i < len(m.word); i++ {
			if i < len(m.input) {
				s += tileStyle.Render(string(m.input[i]))
			} else {
				s += tileStyle.Render(" ")
			}
		}
	} else {
		for i := 0; i < len(m.word); i++ {
			s += tileStyle.Render(" ")
		}
	}
	return s
}

func (m model) EndView() string {
	s := "the word was\n\n"
	s += wordStyle.Foreground(whiteColor).Background(greenColor).Render(string(m.word))
	s += "\n\n"

	if m.guessed {
		try := "try"
		if m.line != 1 {
			try = "tries"
		}
		s += fmt.Sprintf("\n\nYou guessed the word in %d %s.", m.line, try)
	} else {
		s += "You did not guess the word"
	}

	s += "\n\n"
	s += "Play: space\n"
	s += "Quit: ctrl+c"
	return ScreenView(s)
}
