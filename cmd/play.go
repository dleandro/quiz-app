package cmd

import (
    "fmt"
    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "play",
    Short: "Quiz-CLI is a CLI application quiz",
    Long:  `Quiz-CLI is a CLI game with multiple quiz questions`,
    Run: func(cmd *cobra.Command, args []string) {
        askQuestion()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func askQuestion() {
    question := "What is the capital of France?"
    options := []string{"Berlin", "Madrid", "Paris", "Rome"}

    prompt := promptui.Select{
        Label: question,
        Items: options,
    }

    _, result, err := prompt.Run()

    if err != nil {
        fmt.Printf("Prompt failed %v\n", err)
        return
    }

    fmt.Printf("You chose %q\n", result)
}