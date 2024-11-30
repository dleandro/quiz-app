package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	pb "quiz-cli/api/protofiles"

	"google.golang.org/protobuf/proto"
)

func setupTestServer(t *testing.T) (*server, string, string) {
    // Create temp directory for test files
    tmpDir, err := os.MkdirTemp("", "quiz-test-*")
    if err != nil {
        t.Fatalf("failed to create temp dir: %v", err)
    }

    resultsFile = filepath.Join(tmpDir, "results.pb")
    questionsFile = filepath.Join(tmpDir, "questions.pb")

    return &server{}, tmpDir, questionsFile
}

func cleanupTest(tmpDir string) {
    os.RemoveAll(tmpDir)
}

func TestGetQuestions_Empty(t *testing.T) {
    s, tmpDir, _ := setupTestServer(t)
    defer cleanupTest(tmpDir)

    ctx := context.Background()
    res, err := s.GetQuestions(ctx, &pb.Empty{})
    if err != nil {
        t.Fatalf("GetQuestions failed: %v", err)
    }

    if len(res.Questions) != 0 {
        t.Fatalf("Expected empty questions, got %d questions", len(res.Questions))
    }
}

func TestCreateQuestion(t *testing.T) {
    s, tmpDir, _ := setupTestServer(t)
    defer cleanupTest(tmpDir)

    ctx := context.Background()
    req := &pb.CreateQuestionRequest{
        Question: &pb.Question{
            Question: "What is the capital of France?",
            Options:  []string{"Berlin", "Madrid", "Paris", "Rome"},
            Answer:   "Paris",
        },
    }

    _, err := s.CreateQuestion(ctx, req)
    if err != nil {
        t.Fatalf("CreateQuestion failed: %v", err)
    }

    res, err := s.GetQuestions(ctx, &pb.Empty{})
    if err != nil {
        t.Fatalf("GetQuestions failed: %v", err)
    }

    if len(res.Questions) != 1 {
        t.Fatalf("Expected 1 question, got %d", len(res.Questions))
    }

    if res.Questions[0].Question != "What is the capital of France?" {
        t.Fatalf("Expected question to be 'What is the capital of France?', got '%s'", res.Questions[0].Question)
    }
}

func TestDeleteQuestion(t *testing.T) {
    s, tmpDir, _ := setupTestServer(t)
    defer cleanupTest(tmpDir)

    ctx := context.Background()
    createReq := &pb.CreateQuestionRequest{
        Question: &pb.Question{
            Question: "What is the capital of France?",
            Options:  []string{"Berlin", "Madrid", "Paris", "Rome"},
            Answer:   "Paris",
        },
    }

    _, err := s.CreateQuestion(ctx, createReq)
    if err != nil {
        t.Fatalf("CreateQuestion failed: %v", err)
    }

    res, err := s.GetQuestions(ctx, &pb.Empty{})
    if err != nil {
        t.Fatalf("GetQuestions failed: %v", err)
    }

    if len(res.Questions) != 1 {
        t.Fatalf("Expected 1 question, got %d", len(res.Questions))
    }

    deleteReq := &pb.DeleteQuestionRequest{
        Id: res.Questions[0].Id,
    }

    _, err = s.DeleteQuestion(ctx, deleteReq)
    if err != nil {
        t.Fatalf("DeleteQuestion failed: %v", err)
    }

    res, err = s.GetQuestions(ctx, &pb.Empty{})
    if err != nil {
        t.Fatalf("GetQuestions failed: %v", err)
    }

    if len(res.Questions) != 0 {
        t.Fatalf("Expected 0 questions, got %d", len(res.Questions))
    }
}

func TestSaveResults(t *testing.T) {
    s, tmpDir, _ := setupTestServer(t)
    defer cleanupTest(tmpDir)

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

    if len(result.ParticipantResults) != 1 {
        t.Fatalf("Expected 1 participant result, got %d", len(result.ParticipantResults))
    }
}

func TestGetStatistics(t *testing.T) {
    s, tmpDir, _ := setupTestServer(t)
    defer cleanupTest(tmpDir)

    ctx := context.Background()

    // Save some results first
    s.SaveResults(ctx, &pb.ResultsRequest{
        CorrectAnswers: 1,
        TotalQuestions: 5,
    })
    s.SaveResults(ctx, &pb.ResultsRequest{
        CorrectAnswers: 4,
        TotalQuestions: 5,
    })

    req := &pb.ResultsRequest{
		CorrectAnswers: 5,
		TotalQuestions: 5,
	}
    res, err := s.GetStatistics(ctx, req)
    if err != nil {
        t.Fatalf("GetStatistics failed: %v", err)
    }

    if res.PercentageBetterThan <= 0 {
        t.Fatalf("Expected percentage better than 0, got %.2f", res.PercentageBetterThan)
    }
}