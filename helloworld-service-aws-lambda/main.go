package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda" // implements the Lambda programming model for Go
)

type MyEvent struct {
	Name string `json:"name"`
}

// Process events
func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	// The entry point that runs your Lambda function code.
	lambda.Start(HandleRequest)
}
