package db

import (
	"fmt"

	"github.com/djengua/users-go/models"
)

func SignUp(s models.SignUp) error {
	fmt.Println("Begin Registry in DB of new user")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	Db.AutoMigrate(&models.SignUp{})

	newUser := &models.User{}

	newUser.Email = s.UserEmail
	newUser.Uuid = s.UserUUID

	result := Db.Create(newUser)

	msg := fmt.Sprintf("> Rows affected: %v", result.RowsAffected)
	fmt.Println(msg)

	return nil
}
