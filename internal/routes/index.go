package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/app/handlers"
	repository "homestay.com/nguyenduy/internal/app/repositories"
	service "homestay.com/nguyenduy/internal/app/services"
	"homestay.com/nguyenduy/internal/middlewares"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Static("/uploads", "./uploads")
	router.Use(
		middlewares.CORSMiddleware(),
		middlewares.LoggerMiddleware(),
		middlewares.RateLimitingMiddleware())
	go middlewares.CleanUpClients()

	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	hotelRepository := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepository)
	hotelHandler := handlers.NewHotelHandler(hotelService)

	roomRepository := repository.NewRoomRepository(db)
	roomService := service.NewRoomService(roomRepository)
	roomHandler := handlers.NewRoomHandler(roomService)

	roomTypeRepository := repository.NewRoomTypeRepository(db)
	roomTypeService := service.NewRoomTypeService(roomTypeRepository)
	roomTypeHandler := handlers.NewRoomTypeHandler(roomTypeService)

	guestRepository := repository.NewGuestRepository(db)
	guestService := service.NewGuestService(guestRepository)
	guestHandler := handlers.NewGuestHandler(guestService)

	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepository)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	bookingRepository := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepository)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	staffRepository := repository.NewStaffRepository(db, hotelRepository, userRepository)
	staffService := service.NewStaffService(staffRepository)
	staffHandler := handlers.NewStaffHandler(staffService)

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	hotelRoutesPublic := router.Group("/api/hotels")
	{
		hotelRoutesPublic.GET("", hotelHandler.GetAllHotels)
		hotelRoutesPublic.GET(":id", hotelHandler.GetHotelByID)
	}

	roomRoutesPublic := router.Group("/api/rooms")
	{
		roomRoutesPublic.GET("", roomHandler.GetAllRooms)
		roomRoutesPublic.GET(":id", roomHandler.GetRoomByID)
	}

	roomTypeRoutesPublic := router.Group("/api/room-types")
	{
		roomTypeRoutesPublic.GET("", roomTypeHandler.GetAllRoomTypes)
		roomTypeRoutesPublic.GET(":id", roomTypeHandler.GetRoomTypeByID)
	}

	guestRoutesPublic := router.Group("/api/guests")
	{
		guestRoutesPublic.GET("", guestHandler.GetAllGuests)
		guestRoutesPublic.GET(":id", guestHandler.GetGuestByID)
	}

	paymentRoutesPublic := router.Group("/api/payments")
	{
		paymentRoutesPublic.GET("", paymentHandler.GetAllPayments)
		paymentRoutesPublic.GET(":id", paymentHandler.GetPaymentByID)
	}

	bookingRoutesPublic := router.Group("/api/bookings")
	{
		bookingRoutesPublic.GET("", bookingHandler.GetAllBookings)
		bookingRoutesPublic.GET(":id", bookingHandler.GetBookingByID)
	}

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.CheckJwt())
	{
		protectedRoutes.GET("/profile", authHandler.Profile)

		hotelRoutes := protectedRoutes.Group("/hotels")
		{
			hotelRoutes.POST("", hotelHandler.CreateHotel)
			hotelRoutes.PUT(":id", hotelHandler.UpdateHotel)
			hotelRoutes.DELETE(":id", hotelHandler.DeleteHotel)
		}

		roomRoutes := protectedRoutes.Group("/rooms")
		{
			roomRoutes.POST("", roomHandler.CreateRoom)
			roomRoutes.PUT(":id", roomHandler.UpdateRoom)
			roomRoutes.DELETE(":id", roomHandler.DeleteRoom)
		}

		roomTypeRoutes := protectedRoutes.Group("/room-types")
		{
			roomTypeRoutes.POST("", roomTypeHandler.CreateRoomType)
			roomTypeRoutes.PUT(":id", roomTypeHandler.UpdateRoomType)
			roomTypeRoutes.DELETE(":id", roomTypeHandler.DeleteRoomType)
		}

		guestRoutes := protectedRoutes.Group("/guests")
		{
			guestRoutes.POST("", guestHandler.CreateGuest)
			guestRoutes.PUT(":id", guestHandler.UpdateGuest)
			guestRoutes.DELETE(":id", guestHandler.DeleteGuest)
		}

		paymentRoutes := protectedRoutes.Group("/payments")
		{
			paymentRoutes.POST("", paymentHandler.CreatePayment)
			paymentRoutes.PUT(":id", paymentHandler.UpdatePayment)
			paymentRoutes.DELETE(":id", paymentHandler.DeletePayment)
		}

		bookingRoutes := protectedRoutes.Group("/bookings")
		{
			bookingRoutes.POST("", bookingHandler.CreateBooking)
			bookingRoutes.PUT(":id", bookingHandler.UpdateBooking)
			bookingRoutes.DELETE(":id", bookingHandler.DeleteBooking)
		}
	}

	staffRoutes := router.Group("/api/staffs")
	staffRoutes.Use(middlewares.CheckJwt())
	{
		staffRoutes.POST("", staffHandler.CreateStaff)
		staffRoutes.GET("", staffHandler.GetAllStaff)
		staffRoutes.GET(":id", staffHandler.GetStaffByID)
		staffRoutes.PUT(":id", staffHandler.UpdateStaff)
		staffRoutes.DELETE(":id", staffHandler.DeleteStaff)
	}

	return router
}
