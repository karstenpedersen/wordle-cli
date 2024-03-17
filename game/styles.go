package game

import "github.com/charmbracelet/lipgloss"

var whiteColor = lipgloss.Color("#FAFAFA")
var grayColor = lipgloss.AdaptiveColor{Light: "#787c7e", Dark: "#3a3a3c"}
var greenColor = lipgloss.Color("#6aaa64")
var yellowColor = lipgloss.Color("#c9b458")
var backgroundColor = lipgloss.AdaptiveColor{Light: "#FAFAFA", Dark: "#121213"}

var tileStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(whiteColor)).Background(grayColor)
var spaceTileStyle = lipgloss.NewStyle().Background(backgroundColor)
var correctTileStyle = lipgloss.NewStyle().Background(greenColor).Inherit(tileStyle)
var presentTileStyle = lipgloss.NewStyle().Background(yellowColor).Inherit(tileStyle)
var incorrectTileStyle = lipgloss.NewStyle().Inherit(tileStyle)
var cursorTileStyle = lipgloss.NewStyle().Underline(true).Bold(true).Blink(true).Inherit(tileStyle)

var wordStyle = lipgloss.NewStyle().Bold(true)
var indicatorStyle = lipgloss.NewStyle().Blink(true).Foreground(grayColor)

var boardStyle = lipgloss.NewStyle()

var lineBorder = lipgloss.Border{
	Left:  " ",
	Right: " ",
}
var cursorLineBorder = lipgloss.Border{
	Left:  ">",
	Right: "<",
}
var lineBorderStyle = lipgloss.NewStyle().BorderForeground(whiteColor).BorderBackground(backgroundColor)
var lineStyle = lipgloss.NewStyle().Border(lineBorder, false, true).Inherit(lineBorderStyle)
var cursorLineStyle = lipgloss.NewStyle().Border(cursorLineBorder, false, true).Inherit(lineBorderStyle)

var screenStyle = lipgloss.NewStyle().Background(backgroundColor).Align(lipgloss.Center).AlignVertical(lipgloss.Center)
