package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"sort"
	"sync"

	pb "quiz-cli/api/protofiles"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
    "quiz-cli/utils"
)

type server struct {
    pb.UnimplementedQuizServiceServer
    mu sync.Mutex
}

var resultsFile string
var questionsFile string

func main() {

	projectRoot, err := utils.FindProjectRoot()
    if err != nil {
        log.Fatalf("Error finding project root: %v", err)
    }

	dataPath := filepath.Join(projectRoot, "data")

    // Construct the file paths relative to the root of the project
    resultsFile = filepath.Join(dataPath, "results.pb")
    questionsFile = filepath.Join(dataPath, "questions.pb")
	
	// Load environment variables from .env file
    err = godotenv.Load(filepath.Join(projectRoot, ".env"))
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

	port := os.Getenv("API_PORT")
    if port == "" {
        port = "50051" // Default port if API_PORT is not set
    }

    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterQuizServiceServer(s, &server{})
    log.Println("Server is running on port", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func (s *server) GetQuestions(ctx context.Context, in *pb.Empty) (*pb.QuestionsResponse, error) {
    log.Println("A request to get questions has been received")
    var questionsWrapper pb.QuestionsResponse
    if _, err := os.Stat(questionsFile); err == nil {
        data, err := os.ReadFile(questionsFile)
        if err != nil {
            return nil, err
        }
        if err := proto.Unmarshal(data, &questionsWrapper); err != nil {
            return nil, err
        }
    }
    return &pb.QuestionsResponse{Questions: questionsWrapper.Questions}, nil
}

func (s *server) SaveResults(ctx context.Context, in *pb.ResultsRequest) (*pb.ResultsResponse, error) {
    log.Println("A request to save quiz results has been received")
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
    log.Println("A request to get quiz stats has been received")

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

func (s *server) CreateQuestion(ctx context.Context, in *pb.CreateQuestionRequest) (*pb.CreateQuestionResponse, error) {
    log.Println("A request to create a question has been received")
    s.mu.Lock()
    defer s.mu.Unlock()

    var questionsWrapper pb.QuestionsResponse
    if _, err := os.Stat(questionsFile); err == nil {
        data, err := os.ReadFile(questionsFile)
        if err != nil {
            return nil, err
        }
        if err := proto.Unmarshal(data, &questionsWrapper); err != nil {
            return nil, err
        }
    }
    questions := questionsWrapper.Questions

    question := in.Question
    question.Id = uuid.New().String()
    questions = append(questions, question)

    data, err := proto.Marshal(&pb.QuestionsResponse{Questions: questions})
    if err != nil {
        return nil, err
    }
    if err := os.WriteFile(questionsFile, data, 0644); err != nil {
        return nil, err
    }

    return &pb.CreateQuestionResponse{Message: "Question created successfully"}, nil
}

func (s *server) DeleteQuestion(ctx context.Context, in *pb.DeleteQuestionRequest) (*pb.DeleteQuestionResponse, error) {
    log.Println("A request to delete a question has been received")
    s.mu.Lock()
    defer s.mu.Unlock()

    var questionResponse pb.QuestionsResponse
    if _, err := os.Stat(questionsFile); err == nil {
        data, err := os.ReadFile(questionsFile)
        if err != nil {
            return nil, err
        }
        if err := proto.Unmarshal(data, &questionResponse); err != nil {
            return nil, err
        }
    }

	var questions = questionResponse.Questions

    for i, q := range questions {
        if q.Id == in.Id {
            questions = append(questions[:i], questions[i+1:]...)
            break
        }
    }

    data, err := proto.Marshal(&pb.QuestionsResponse{Questions: questions})
    if err != nil {
        return nil, err
    }
    if err := os.WriteFile(questionsFile, data, 0644); err != nil {
        return nil, err
    }

    return &pb.DeleteQuestionResponse{Message: "Question deleted successfully"}, nil
}

