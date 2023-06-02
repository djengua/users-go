package db

import (
	"fmt"
	"os"

	"github.com/djengua/users-go/models"
	"github.com/djengua/users-go/secretmgr"
	"github.com/jinzhu/gorm"
)

var SecretModel models.SecretRDS
var err error
var Db *gorm.DB

func ReadSecret() error {
	SecretModel, err = secretmgr.GetSecret(os.Getenv("SECRET_NAME"))
	return err
}

func DbConnect() error {
	fmt.Println("Try connect to Database...")
	connector := PostgresConnector{}
	Db, err = connector.GetConnection()

	if err != nil {
		fmt.Println(err.Error())
		// panic("Database connection attempt was unsuccessful...")
		return err
	}

	err = Db.DB().Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer Db.Close()

	fmt.Println("Database Connected successfully.....")

	return nil
}
