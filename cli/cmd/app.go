package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "quiz-cli/api/protofiles"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiClient struct {
	conn *grpc.ClientConn
	c    pb.QuizServiceClient
	ctx  context.Context
}

var apiConn *ApiClient

var rootCmd = &cobra.Command{
	Use:   "quiz-cli",
	Short: "Quiz-CLI is a CLI application quiz",
	Long:  `Quiz-CLI is a CLI game with multiple quiz questions`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		apiURL := os.Getenv("API_URL")
		apiPort := os.Getenv("API_PORT")
		if apiURL == "" || apiPort == "" {
			log.Fatalf("Mandatory environment variables are not set. Check .env file for API_URL and API_PORT")
		}

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", apiURL, apiPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to the API: %v", err)
		}

		apiConn = &ApiClient{
			conn: conn,
			c:    pb.NewQuizServiceClient(conn),
			ctx:  context.Background(),
		}

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Close the connection
		if apiConn.conn != nil {
			apiConn.conn.Close()
		}
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
	Run: func(cmd *cobra.Command, args []string) {
		runQuiz()
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get currently available questions",
	Run: func(cmd *cobra.Command, args []string) {
		getQuestions()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new question",
	Run: func(cmd *cobra.Command, args []string) {
		createQuestion()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a question",
	Run: func(cmd *cobra.Command, args []string) {
		deleteQuestion()
	},
}

func runQuiz() {
	ctx := apiConn.ctx
	c := apiConn.c

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

	stats, err := c.GetStatistics(ctx, &pb.ResultsRequest{
		CorrectAnswers: int32(correctAnswers),
		TotalQuestions: int32(len(r.Questions)),
	})
	if err != nil {
		log.Fatalf("could not get statistics: %v", err)
	}

	_, err = c.SaveResults(ctx, &pb.ResultsRequest{
		CorrectAnswers: int32(correctAnswers),
		TotalQuestions: int32(len(r.Questions)),
	})
	if err != nil {
		log.Fatalf("could not save results: %v", err)
	}

	log.Printf("You answered %d out of %d questions correctly.\n", correctAnswers, len(r.Questions))
	log.Printf("You were better than %.f%% of all quizzers.\n", stats.PercentageBetterThan)
}

func getQuestions() {
	c := apiConn.c
	ctx := apiConn.ctx

	// Fetch available questions
	r, err := c.GetQuestions(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get questions: %v", err)
	}

	if len(r.Questions) == 0 {
		log.Println("No questions available")
		return
	}

	// Display questions
	for i, q := range r.Questions {
		log.Printf("\nQuestion %d: %s\n", i+1, q.Question)
		log.Println("Options:")
		for j, opt := range q.Options {
			log.Printf("  %d. %s\n", j+1, opt)
		}
	}
}

func createQuestion() {
	c := apiConn.c
	ctx := apiConn.ctx
	question := promptForQuestion()

	_, err := c.CreateQuestion(ctx, &pb.CreateQuestionRequest{
		Question: &question,
	})
	if err != nil {
		log.Fatalf("could not create question: %v", err)
	}

	log.Println("Question created successfully")
}

func deleteQuestion() {
	c := apiConn.c
	ctx := apiConn.ctx

	// Fetch available questions
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

func promptForQuestionID() string {
	prompt := promptui.Prompt{
		Label: "Question ID",
	}
	questionID, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return questionID
}
