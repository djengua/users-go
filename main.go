package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	// "github.com/djengua/users-go/awsgo"
	"github.com/djengua/users-go/models"
)

func HandleRequest(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	// awsgo.AWSInit()

	// if !validateParams() {
	// 	fmt.Println("'SecretManager' not provided in parameters.")
	// 	err := errors.New("'SecretManager' not provided in parameters")
	// 	return event, err
	// }

	dataSignUp := models.SignUp{}

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			dataSignUp.UserEmail = att
		case "sub":

			dataSignUp.UserUUID = att
		}
	}

	// err := db.ReadSecret()
	// if err != nil {
	// 	fmt.Println("Error reading the secret " + err.Error())
	// 	return event, err
	// }

	return event, nil
}

func validateParams() bool {
	var hasValue bool
	_, hasValue = os.LookupEnv("SecretName")
	return hasValue
}

func main() {
	lambda.Start(HandleRequest)
}
