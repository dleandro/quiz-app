package api

import (
    "context"
    "os"
    "testing"

    "google.golang.org/protobuf/proto"
    pb "quiz-cli/api/protofiles"
)

func setupTestServer() *server {
    return &server{}
}

func TestGetQuestions(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()
    req := &pb.Empty{}

    res, err := s.GetQuestions(ctx, req)
    if err != nil {
        t.Fatalf("GetQuestions failed: %v", err)
    }

    if len(res.Questions) != 2 {
        t.Fatalf("Expected 2 questions, got %d", len(res.Questions))
    }
}

func TestSaveResults(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()
    req := &pb.ResultsRequest{
        CorrectAnswers: 3,
        TotalQuestions: 5,
    }

    _, err := s.SaveResults(ctx, req)
    if err != nil {
        t.Fatalf("SaveResults failed: %v", err)
    }

    var result pb.Result
    data, err := os.ReadFile(resultsFile)
    if err != nil {
        t.Fatalf("Failed to read results file: %v", err)
    }
    if err := proto.Unmarshal(data, &result); err != nil {
        t.Fatalf("Failed to unmarshal results: %v", err)
    }

    if len(result.ParticipantResults) < 1 {
        t.Fatalf("Expected at least 1 participant result, got %d", len(result.ParticipantResults))
    }
}

func TestGetStatistics(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    // Save some results first
    s.SaveResults(ctx, &pb.ResultsRequest{
        CorrectAnswers: 3,
        TotalQuestions: 5,
    })
    s.SaveResults(ctx, &pb.ResultsRequest{
        CorrectAnswers: 4,
        TotalQuestions: 5,
    })

    req := &pb.ResultsRequest{
        CorrectAnswers: 0,
        TotalQuestions: 5,
    }
    res, err := s.GetStatistics(ctx, req)
    if err != nil {
        t.Fatalf("GetStatistics failed: %v", err)
    }

    if res.PercentageBetterThan > 0 {
        t.Fatalf("Expected percentage equal to 0, got %.2f", res.PercentageBetterThan)
    }
}