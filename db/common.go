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

func ReadSecret() error {
	SecretModel, err = secretmgr.GetSecret(os.Getenv("SECRET_NAME"))
	return err
}

func DbConnect() (*gorm.DB, error) {
	fmt.Println("Try connect to Database...")
	connector := PostgresConnector{}
	db, err := connector.GetConnection()

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Database connection attempt was unsuccessful...")
		return nil, err
	}

	// Ping a la base de datos
	err = db.DB().Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Database Connected successfully.....")

	return db, nil
}
