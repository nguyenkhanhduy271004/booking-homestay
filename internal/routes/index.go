package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/app/handlers"
	"homestay.com/nguyenduy/internal/app/repository"
	"homestay.com/nguyenduy/internal/app/service"
	"homestay.com/nguyenduy/internal/middlewares"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	hotelRepository := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepository)
	hotelHandler := handlers.NewHotelHandler(hotelService)

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	hotelRoutesPublic := router.Group("/api/hotels")
	{
		hotelRoutesPublic.GET("", hotelHandler.GetAll)
		hotelRoutesPublic.GET("/:id", hotelHandler.GetByID)
	}

	protectedRoutes := router.Group("api")
	protectedRoutes.Use(middlewares.CheckJwt())
	{
		protectedRoutes.GET("/profile", authHandler.Profile)

		hotelRoutes := protectedRoutes.Group("/hotels")
		{
			hotelRoutes.POST("", hotelHandler.Create)
			hotelRoutes.PUT("/:id", hotelHandler.Update)
			hotelRoutes.DELETE("/:id", hotelHandler.Delete)
		}
	}

	return router
}
