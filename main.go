package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/djengua/users-go/db"
	"github.com/djengua/users-go/models"
)

func HandleRequest(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	// awsgo.AWSInit()

	// Validar que los datos no esten vacios
	// if !validateParams() {
	// 	fmt.Println("'SecretManager' not provided in parameters.")
	// 	err := errors.New("'SecretManager' not provided in parameters")
	// 	return event, err
	// }

	dataSignUp := models.SignUp{}

	fmt.Println(" > Reading data from cognito... ")
	fmt.Println(event.Request.UserAttributes)
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			dataSignUp.UserEmail = att
		case "sub":
			dataSignUp.UserUUID = att
		}
	}

	// Leer información del secret para la BD
	// err := db.ReadSecret()
	// if err != nil {
	// 	fmt.Println("Error reading the secret " + err.Error())
	// 	return event, err
	// }

	err := db.SignUp(dataSignUp)
	if err != nil {
		fmt.Println("Error " + err.Error())
		return event, err
	}

	return event, nil
}

func validateParams() bool {
	var hasValue bool
	_, hasValue = os.LookupEnv("SECRET_NAME")
	return hasValue
}

func main() {
	lambda.Start(HandleRequest)
}
