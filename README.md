# Quiz CLI

The Quiz CLI is a command-line application that allows users to take a quiz and see how they compare to other users. The application consists of two main components: the API server and the CLI client. The API server provides the quiz questions and stores the results, while the CLI client interacts with the user to take the quiz.

## Architecture
The architecture of the Quiz CLI application is composed of the following components:

1. **API Server**: The API server is built using Go and gRPC. It provides endpoints for fetching quiz questions, saving quiz results, and retrieving statistics about user performance. The server stores quiz results in a file (`results.pb`) using Protocol Buffers for serialization.

2. **CLI Client**: The CLI client is a command-line application built using Go. It interacts with the user to take the quiz, submits the results to the API server, and displays statistics about the user's performance compared to other users. The CLI client communicates with the API server using gRPC.

## Features

- Take a quiz with multiple-choice questions
- Save quiz results
- Compare your performance with other users

## Prerequisites

- Docker (optional)
- Go

## Running the Application

1. **Clone the repository**:
    ```sh
    git clone https://github.com/yourusername/quiz-cli.git
    cd quiz-cli
    ```

2. **Create a `.env` file**:
    ```sh
    cp .env.example .env
    ```

### Running the api with Docker Compose

**Build and run the container**:
```sh
docker-compose up --build
```

#### Or locally

```sh
go run ./api/server.go
```

### Running the CLI

```sh
go run ./cli/app.go
```

### Run the tests

```sh
go run ./api
```

## Important notes

The API should always be ran first or else the cli won't work due to not having connection. We could solve this with polling for connection in the future.

Also to reset the results from past executions you can delete the results.pb file in the root of the repo

Next steps:

- Separate each endpoint in its own controller file
- Refactor the cli code in order to be more testable for example by sending results in the end instead of logging
- Could have adapters for the dependencies
- Dependency injection
- Extract the Parsing and reading of proto files to infra adapters or aux methods (basically extracting duplicated code)
- Error handling 