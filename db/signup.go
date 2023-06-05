package db

import (
	"fmt"

	"github.com/djengua/users-go/models"
)

// Registro d enuevo usuario, proveniente de cognito
func SignUp(s models.SignUp) error {
	fmt.Println("Begin Registry in DB of new user")
	db, err := DbConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	result := db.AutoMigrate(&models.User{})
	newUser := &models.User{}
	newUser.Email = s.UserEmail
	newUser.Uuid = s.UserUUID

	result = db.Create(newUser)

	msg := fmt.Sprintf("> Rows affected: %v", result.RowsAffected)
	fmt.Println(msg)

	return nil
}
