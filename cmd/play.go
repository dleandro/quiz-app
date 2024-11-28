package cmd

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/manifoldco/promptui"
    "github.com/spf13/cobra"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "quiz-cli/api/protofiles"
)

var rootCmd = &cobra.Command{
    Use:   "play",
    Short: "Quiz-CLI is a CLI application quiz",
    Long:  `Quiz-CLI is a CLI game with multiple quiz questions`,
    Run: func(cmd *cobra.Command, args []string) {
        runQuiz()
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func runQuiz() {
    conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewQuizServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.GetQuestions(ctx, &pb.Empty{})
    if err != nil {
        log.Fatalf("could not get questions: %v", err)
    }

    for _, q := range r.Questions {
        prompt := promptui.Select{
            Label: q.Question,
            Items: q.Options,
        }

        _, result, err := prompt.Run()

        if err != nil {
            fmt.Printf("Prompt failed %v\n", err)
            return
        }

        if result == q.Answer {
            fmt.Printf("Correct! %q is the right answer.\n", result)
        } else {
            fmt.Printf("Wrong! The correct answer is %q.\n", q.Answer)
        }
    }
}