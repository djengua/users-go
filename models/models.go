package models

import "github.com/jinzhu/gorm"

type SecretRDS struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier int    `json:"dbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}

type User struct {
	gorm.Model
	Uuid  string `json:"uuid"`
	Email string `json:"email" gorm:index`
}
