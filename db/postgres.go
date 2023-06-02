package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type PostgresConnector struct {
}

func (p *PostgresConnector) GetConnection() (db *gorm.DB, err error) {
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	dbUser := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPort, password)

	return gorm.Open("postgres", dbURI)
}
