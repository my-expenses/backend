package routes

import (
	usersRoutes "backend/routes/users"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	users := r.Group("users")
	usersRoutes.InitializeUsersRoutes(users)
}
