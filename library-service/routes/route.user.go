package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	auth_controller "library-service/controllers/auth-controller"
	auth_handler "library-service/handlers/auth-handler"
)

func InitUserRoutes(db *gorm.DB, route *gin.RouterGroup) {

	userRepository := auth_controller.NewUserRepository(db)
	userService := auth_controller.NewUserService(userRepository)
	userHandler := auth_handler.NewAuthHandler(userService)

	route.POST("/register", userHandler.RegisterHandler)
	route.POST("/login", userHandler.LoginHandler)
}
