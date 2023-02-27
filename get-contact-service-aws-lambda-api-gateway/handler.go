package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
)

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received req %#v", req)

	switch req.HTTPMethod {
	case "GET":
		return processGet(ctx, req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func processGet(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, _ := req.PathParameters["id"]
	return processGetContact(ctx, id)
}

func processGetContact(ctx context.Context, id string) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received GET Contact request with id = %s", id)

	contactRequest, err := getItem(ctx, id)
	if err != nil {
		return serverError(err)
	}

	if contactRequest == nil {
		return clientError(http.StatusNotFound)
	}

	json, err := json.Marshal(contactRequest)
	if err != nil {
		return serverError(err)
	}
	log.Printf("Successfully fetched Contact item %s", json)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(status),
		StatusCode: status,
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
	}, nil
}
