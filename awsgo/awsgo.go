package awsgo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

// Cargar la configuracion inicial
func AWSInit() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	fmt.Println(" > Config loaded OK ...")

	if err != nil {
		panic("Error at loading configuration .aws/config" + err.Error())
	}
}
