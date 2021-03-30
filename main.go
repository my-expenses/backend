package main

import (
	db "backend/database"
	"backend/routes"
	"backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	customValidator := &utils.CustomValidator{Validator: validator.New()}
	customValidator.TranslateErrors()
	e.Validator = customValidator

	utils.InitializeEnvVars()
	db.MigrateTables()
	routes.InitializeRoutes(e)

	e.Start(":3000")
}
