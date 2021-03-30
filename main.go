package main

import (
	db "backend/database"
	"backend/routes"
	"backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()

	customValidator := &utils.CustomValidator{Validator: validator.New()}
	customValidator.TranslateErrors()
	e.Validator = customValidator

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
	db.MigrateTables()
	routes.InitializeRoutes(e)

	e.Start(":3000")
}
