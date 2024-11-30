package main

import (
	"log"
	"os"

	"quiz-cli/cli/cmd"
	"quiz-cli/cli/infrastructure"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Execute()
}

var apiClient *infrastructure.ApiClient
var rootCmd = &cobra.Command{
	Use:   "quiz-cli",
	Short: "Quiz-CLI is a CLI application quiz",
	Long:  `Quiz-CLI is a CLI game with multiple quiz questions`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		apiClient, err = infrastructure.NewApiClient()
		if err != nil {
			log.Fatalf("Failed to create API client: %v", err)
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		infrastructure.CloseConnection(apiClient)
	},
}

func Execute() {
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(getCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play the quiz",
	Run: func(_ *cobra.Command, args []string) {
		cmd.PlayQuiz(apiClient)
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get currently available questions",
	Run: func(_ *cobra.Command, args []string) {
		cmd.GetQuestions(apiClient)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new question",
	Run: func(_ *cobra.Command, args []string) {
		cmd.CreateQuestion(apiClient)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a question",
	Run: func(_ *cobra.Command, args []string) {
		cmd.DeleteQuestion(apiClient)
	},
}
