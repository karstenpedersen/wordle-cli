package cmd

import (
	"github.com/karstenpedersen/wordle-cli/game"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wordle",
	Short: "Play Wordle",
	Long:  "Play Wordle",
	Example: `Make a friend guess a custom word:
  wordle --word cool`,
	Version: "0.1.0",
	Aliases: []string{"wordle"},
	Run: func(cmd *cobra.Command, args []string) {
		word, _ := cmd.Flags().GetString("word")
    wordlistPath, _ := cmd.Flags().GetString("wordlist")
		game.Start(word, wordlistPath)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("word", "w", "", "provide a word to guess")
	rootCmd.Flags().StringP("wordlist", "W", "./wordlists/wordle", "provide a custom wordlist")
}
