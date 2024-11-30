package cmd

import (
	"log"
	pb "quiz-cli/api/protofiles"
	"quiz-cli/cli/infrastructure"
)

func GetQuestions(apiClient *infrastructure.ApiClient) {
	c := apiClient.C
	ctx := apiClient.Ctx

	r, err := c.GetQuestions(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get questions: %v", err)
	}

	displayQuestions(r.Questions)

}

func displayQuestions(questions []*pb.Question) {
	if len(questions) == 0 {
		log.Println("No questions available")
		return
	}

	for i, q := range questions {
		log.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		log.Println("Options:")
		for j, opt := range q.Options {
			log.Printf("  %d. %s\n", j+1, opt)
		}
	}
}
