package users

import (
	usersController "backend/controllers/users"
	"github.com/gin-gonic/gin"
)

func InitializeUsersRoutes(usersRoute *gin.RouterGroup) {
	usersRoute.POST("/register", usersController.NewUser)
}

