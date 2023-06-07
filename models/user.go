package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Uuid     string `json:"uuid"`
	Email    string `json:"email" omitempty slq:index`
	Names    string `json:"names"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
	Type     int    `json:"type"`
	Active   bool   `json:"active"`
}
