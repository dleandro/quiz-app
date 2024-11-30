package cmd

import (
	"log"
	pb "quiz-cli/api/protofiles"

	"quiz-cli/cli/infrastructure"
	"github.com/manifoldco/promptui"

)

func PlayQuiz(apiClient *infrastructure.ApiClient) {
	ctx := apiClient.Ctx
	c := apiClient.C

	r, err := c.GetQuestions(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get questions: %v", err)
	}

	correctAnswers := 0

	for _, q := range r.Questions {
		prompt := promptui.Select{
			Label: q.Question,
			Items: q.Options,
		}

		_, result, err := prompt.Run()

		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}

		if result == q.Answer {
			correctAnswers++
		}
	}

	showResults(apiClient, r, correctAnswers)
	
	_, err = c.SaveResults(ctx, &pb.ResultsRequest{
		CorrectAnswers: int32(correctAnswers),
		TotalQuestions: int32(len(r.Questions)),
	})
	if err != nil {
		log.Fatalf("Could not save results: %v", err)
	}
}

func showResults(apiClient *infrastructure.ApiClient, r *pb.QuestionsResponse, correctAnswers int) {
	stats, err := apiClient.C.GetStatistics(apiClient.Ctx, &pb.ResultsRequest{
		CorrectAnswers: int32(correctAnswers),
		TotalQuestions: int32(len(r.Questions)),
	})
	if err != nil {
		log.Fatalf("Could not get statistics: %v", err)
	}

	log.Printf("You answered %d out of %d questions correctly.\n", correctAnswers, len(r.Questions))
	log.Printf("You were better than %.f%% of all quizzers.\n", stats.PercentageBetterThan)
}