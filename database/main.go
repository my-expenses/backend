package database

import (
	usersModel "backend/models/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	databaseConnection *gorm.DB = nil
	err                error
)

func MigrateTables() {
	initializeDBConnection()
	if err == nil {
		databaseConnection.AutoMigrate(
			&usersModel.User{},
			&usersModel.Credential{},
		)
	}
}

func initializeDBConnection() {
	dsn := os.ExpandEnv("${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8mb4&parseTime=True&loc=Africa%2FCairo")
	connection, connectionErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if connectionErr == nil {
		databaseConnection = connection
	}
	err = connectionErr
}

func GetDBConnection() *gorm.DB {
	if databaseConnection == nil {
		initializeDBConnection()
	}
	return databaseConnection
}
