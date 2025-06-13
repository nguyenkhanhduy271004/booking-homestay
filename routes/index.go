package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"homestay.com/nguyenduy/app/controller"
	"homestay.com/nguyenduy/app/repository"
	"homestay.com/nguyenduy/app/service"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	return router
}
