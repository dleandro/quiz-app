package main

import (
	"log"
	"quiz-cli/cli/cmd"

	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    cmd.Execute()
}

// optional 
// addQuestionWithAnswers