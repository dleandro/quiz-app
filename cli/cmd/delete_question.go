package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	pb "quiz-cli/api/protofiles"
	"quiz-cli/cli/infrastructure"
)

func DeleteQuestion(apiClient *infrastructure.ApiClient) {
	c := apiClient.C
	ctx := apiClient.Ctx

	r, err := c.GetQuestions(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get questions: %v", err)
	}

	if len(r.Questions) == 0 {
		log.Println("No questions available to delete")
		return
	}

	// Create list of questions for selection
	var items []string
	questionMap := make(map[string]string) // map display string to ID

	for _, q := range r.Questions {
		displayString := fmt.Sprintf("%s", q.Question)
		items = append(items, displayString)
		questionMap[displayString] = q.Id
	}

	// Prompt user to select question
	prompt := promptui.Select{
		Label: "Select question to delete",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	// Get ID of selected question
	questionID := questionMap[result]

	// Call delete with selected ID
	_, err = c.DeleteQuestion(ctx, &pb.DeleteQuestionRequest{
		Id: questionID,
	})
	if err != nil {
		log.Fatalf("could not delete question: %v", err)
	}

	log.Printf("Question '%s' deleted successfully\n", result)
}
