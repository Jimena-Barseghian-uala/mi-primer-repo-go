package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

func handleRequest(ctx context.Context, e events.DynamoDBEvent) {
	topicARN := "arn:aws:sns:us-east-1:645342462303:ContactTopic"
	awsRegion := "us-east-1"

	awsSession, _ := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})

	snsService := sns.New(awsSession)

	for _, record := range e.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		id := record.Change.NewImage["id"].String()
		firstName := record.Change.NewImage["firstName"].String()
		lastName := record.Change.NewImage["lastName"].String()
		status := record.Change.NewImage["status"].String()

		params := &sns.PublishInput{
			Message:  aws.String("Id: " + id + "\nFirst Name: " + firstName + "\nLast Name: " + lastName + "\nStatus: " + status + "\nSuccessful SNS Email sent from Contacts Table."),
			TopicArn: aws.String(topicARN),
		}

		resp, err := snsService.Publish(params)
		if err != nil {
			log.Printf("error from call to snsService.Publish: %v", err)
		}

		log.Printf("response: %v", resp)
	}
}

func main() {
	lambda.Start(handleRequest)
}
