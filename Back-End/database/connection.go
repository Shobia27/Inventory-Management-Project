package database

import (
	"fmt"

	"myproject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	port     = 5432
	dbname   = "inventory_mgmt"
)

var DB *gorm.DB

func Connect() {
	userBDInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	connection, err := gorm.Open(postgres.Open(userBDInfo), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{}, &models.Product{}, &models.Supplier{})
}
