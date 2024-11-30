package cmd

import (
	"log"

	"fmt"
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	pb "quiz-cli/api/protofiles"
	"quiz-cli/cli/infrastructure"
)

func CreateQuestion(apiClient *infrastructure.ApiClient) {
	c := apiClient.C
	ctx := apiClient.Ctx
	question := promptForQuestion()

	_, err := c.CreateQuestion(ctx, &pb.CreateQuestionRequest{
		Question: &question,
	})
	if err != nil {
		log.Fatalf("could not create question: %v", err)
	}

	log.Println("Question created successfully")
}

func promptForQuestion() pb.Question {
	prompt := promptui.Prompt{
		Label: "Question",
	}
	questionText, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	options := []string{}
	for i := 0; i < 4; i++ {
		prompt := promptui.Prompt{
			Label: fmt.Sprintf("Option %d", i+1),
		}
		option, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
		}
		options = append(options, option)
	}

	prompt = promptui.Prompt{
		Label: "Answer",
	}
	answer, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return pb.Question{
		Id:       uuid.New().String(),
		Question: questionText,
		Options:  options,
		Answer:   answer,
	}
}
