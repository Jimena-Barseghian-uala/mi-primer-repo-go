package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"log"
)

type CreateContact struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

var validate *validator.Validate = validator.New()

func processPost(ctx context.Context, createContact CreateContact) (*Contact, error) {
	err := validate.Struct(&createContact)
	if err != nil {
		log.Printf("Invalid body: %v", err)
		log.Println(err.Error())
		return nil, err
	}
	log.Printf("Received POST request with item: %+v", createContact)

	res, err := insertItem(ctx, createContact)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Printf("Inserted new contact: %+v", res)

	return res, nil
}
