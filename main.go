package main

import (
	db "backend/database"
	"backend/routes"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	utils.InitializeEnvVars()
	db.MigrateTables()
	routes.InitializeRoutes(r)

	r.Run(":3000")
}
