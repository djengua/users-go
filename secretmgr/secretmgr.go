package secretmgr

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/djengua/users-go/awsgo"
	"github.com/djengua/users-go/models"
)

func GetSecret(name string) (models.SecretRDS, error) {
	var data models.SecretRDS
	fmt.Println(" > Get Secret" + name)
	service := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := service.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	})

	if err != nil {
		fmt.Println(err.Error())
		return data, err
	}

	json.Unmarshal([]byte(*key.SecretString), &data)
	fmt.Println(" > Reading Secret OK " + name)

	return data, nil
}
