package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"sync"

	pb "quiz-cli/api/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
    pb.UnimplementedQuizServiceServer
    mu sync.Mutex
}

const resultsFile = "results.pb"

func (s *server) GetQuestions(ctx context.Context, in *pb.Empty) (*pb.QuestionsResponse, error) {
    fmt.Println("A request to get questions has been received")
    questions := []*pb.Question{
        {Question: "What is the capital of France?", Options: []string{"Berlin", "Madrid", "Paris", "Rome"}, Answer: "Paris"},
        {Question: "What is the capital of Germany?", Options: []string{"Berlin", "Madrid", "Paris", "Rome"}, Answer: "Berlin"},
    }
    return &pb.QuestionsResponse{Questions: questions}, nil
}

func (s *server) SaveResults(ctx context.Context, in *pb.ResultsRequest) (*pb.ResultsResponse, error) {
    fmt.Println("A request to save quiz results has been received")
    s.mu.Lock()
    defer s.mu.Unlock()

    var result pb.Result
    if _, err := os.Stat(resultsFile); err == nil {
        data, err := os.ReadFile(resultsFile)
        if err != nil {
            return nil, err
        }
        if err := proto.Unmarshal(data, &result); err != nil {
            return nil, err
        }
    }

     participantResult := &pb.ParticipantResult{
        CorrectAnswers: in.CorrectAnswers,
        TotalQuestions: in.TotalQuestions,
    }
    result.ParticipantResults = append(result.ParticipantResults, participantResult)


    data, err := proto.Marshal(&result)
    if err != nil {
        return nil, err
    }
    if err := os.WriteFile(resultsFile, data, 0644); err != nil {
        return nil, err
    }

    return &pb.ResultsResponse{Message: "Results saved successfully"}, nil
}

func (s *server) GetStatistics(ctx context.Context, in *pb.ResultsRequest) (*pb.StatisticsResponse, error) {
    fmt.Println("A request to get quiz stats has been received")

    s.mu.Lock()
    defer s.mu.Unlock()

    var result pb.Result
    if _, err := os.Stat(resultsFile); err == nil {
        data, err := os.ReadFile(resultsFile)
        if err != nil {
            return nil, err
        }
        if err := proto.Unmarshal(data, &result); err != nil {
            return nil, err
        }
    }

	    if len(result.ParticipantResults) == 0 {
        return &pb.StatisticsResponse{PercentageBetterThan: 0}, nil
    }

    scores := make([]float32, len(result.ParticipantResults))
    for i, pr := range result.ParticipantResults {
        scores[i] = float32(pr.CorrectAnswers / pr.TotalQuestions)
    }
    sort.Slice(scores, func(i, j int) bool {
        return scores[i] < scores[j]
    })

    currentScore := float32(in.CorrectAnswers / in.TotalQuestions)
    betterThanCount := 0
    for _, score := range scores {
        if currentScore > score {
            betterThanCount++
        }
    }

    percentageBetterThan := (float32(betterThanCount) / float32(len(scores))) * 100

    return &pb.StatisticsResponse{PercentageBetterThan: percentageBetterThan}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterQuizServiceServer(s, &server{})
    fmt.Println("Server is running on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}