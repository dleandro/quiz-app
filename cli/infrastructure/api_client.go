package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "quiz-cli/api/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type ApiClient struct {
	Conn *grpc.ClientConn
	C    pb.QuizServiceClient
	Ctx  context.Context
}

func NewApiClient() (*ApiClient, error) {
	apiURL := os.Getenv("API_URL")
		apiPort := os.Getenv("API_PORT")
		if apiURL == "" || apiPort == "" {
			log.Fatalf("Mandatory environment variables are not set. Check .env file for API_URL and API_PORT")
		}

		conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", apiURL, apiPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to the API: %v", err)
		}

		return &ApiClient{
			Conn: conn,
			C:    pb.NewQuizServiceClient(conn),
			Ctx:  context.Background(),
		}, nil

}

func CloseConnection(c *ApiClient) {
	if c.Conn != nil {
		c.Conn.Close()
	}
}