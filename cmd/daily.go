package cmd

import (
	"encoding/json"
	"github.com/karstenpedersen/wordle-cli/game"
	"github.com/spf13/cobra"
	"log"
	"net/http"
  "fmt"
  "time"
)

var dailyCmd = &cobra.Command{
	Use:   "daily",
	Short: "Play Wordle",
	Long:  "Play Wordle",
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
    if !cmd.Flags().Changed("date") {
      date = time.Now().Format("2006-01-02")
    }
		wordleResponse := getWordleWord(date)
		game.Start(wordleResponse.Solution, "")
	},
}

func init() {
	dailyCmd.Flags().StringP("date", "d", "", "YYYY-MM-DD")
	rootCmd.AddCommand(dailyCmd)
}

type WordleResponse struct {
	Id              int    `json:"id"`
	Solution        string `json:"solution"`
	PrintDate       string `json:"print_date"`
	DaysSinceLaunch int    `json:"days_since_launch"`
	Editor          string `json:"editor"`
}

func getWordleWord(date string) WordleResponse {
	endpoint := fmt.Sprintf("https://www.nytimes.com/svc/wordle/v2/%s.json", date)
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Error:", res.Status)
	}

	var wordleResponse WordleResponse
	err = json.NewDecoder(res.Body).Decode(&wordleResponse)
  if err != nil {
    log.Fatal("Error:", err)
  }
	return wordleResponse
}
