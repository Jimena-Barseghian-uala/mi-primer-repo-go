package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"regexp"
)

func getIdContact(str string) string {
	r := regexp.MustCompile(`Id:\s*(.*?)\s*First Name:`)
	matches := r.FindAllStringSubmatch(str, -1)
	var id string
	for _, v := range matches {
		id = v[1]
	}
	return id
}

func handlerRequest(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		fmt.Printf("Processing request data for event ID %s.\n", snsRecord.MessageAttributes)
		fmt.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		fmt.Printf("SNS Message %s", snsRecord.Message)

		id := getIdContact(snsRecord.Message)
		updateItem(ctx, id)
		fmt.Printf("Contact Updated!")
	}
}

func main() {
	lambda.Start(handlerRequest)
}
