package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"log"
)

const TableName = "Contacts"

var db dynamodb.Client

func init() {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	db = *dynamodb.NewFromConfig(sdkConfig)
}

type Contact struct {
	Id        string `json:"id" dynamodbav:"id"`
	FirstName string `json:"firstName" dynamodbav:"firstName"`
	LastName  string `json:"lastName" dynamodbav:"lastName"`
	Status    string `json:"status" dynamodbav:"status"`
}

func insertItem(ctx context.Context, createContact CreateContact) (*Contact, error) {
	contact := Contact{
		FirstName: createContact.FirstName,
		LastName:  createContact.LastName,
		Status:    "CREATED",
		Id:        uuid.NewString(),
	}

	item, err := attributevalue.MarshalMap(contact)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      item,
	}

	res, err := db.PutItem(ctx, input)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(res.Attributes, &contact)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}
