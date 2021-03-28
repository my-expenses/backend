package main

import (
	db "backend/database"
	"backend/routes"
	"backend/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	utils.InitializeEnvVars()
	db.MigrateTables()
	routes.InitializeRoutes(e)

	e.Start(":3000")
}
