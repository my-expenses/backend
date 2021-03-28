package routes

import (
	usersRoutes "backend/routes/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	CORSMiddleware(r)
	users := r.Group("users")
	usersRoutes.InitializeUsersRoutes(users)
}

func CORSMiddleware(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://192.168.1.100:8080",
		},
		AllowMethods: []string{"POST, OPTIONS, GET, PUT, DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
	}))
}
