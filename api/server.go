package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    pb "quiz-cli/api/protofiles"
)

type server struct {
    pb.UnimplementedQuizServiceServer
}

func (s *server) GetQuestions(ctx context.Context, in *pb.Empty) (*pb.QuestionsResponse, error) {
    questions := []*pb.Question{
        {Question: "What is the capital of France?", Options: []string{"Berlin", "Madrid", "Paris", "Rome"}, Answer: "Paris"},
        {Question: "What is the capital of Germany?", Options: []string{"Berlin", "Madrid", "Paris", "Rome"}, Answer: "Berlin"},
    }
    return &pb.QuestionsResponse{Questions: questions}, nil
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