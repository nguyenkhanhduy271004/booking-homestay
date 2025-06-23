package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"homestay.com/nguyenduy/internal/app/handlers"
	"homestay.com/nguyenduy/internal/app/repositories"
	"homestay.com/nguyenduy/internal/app/services"
	"homestay.com/nguyenduy/internal/middlewares"
)

func InitRoute(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Static("/uploads", "./uploads")

	router.Use(
		middlewares.CORSMiddleware(),
		middlewares.LoggerMiddleware(),
		middlewares.ApiKeyMiddleware(),
		middlewares.RateLimitingMiddleware(),
	)
	go middlewares.CleanUpClients()

	authHandler := handlers.NewAuthHandler(
		services.NewAuthService(repositories.NewUserRepository(db), repositories.NewRoleRepository(db)),
	)

	hotelHandler := handlers.NewHotelHandler(
		services.NewHotelService(repositories.NewHotelRepository(db)),
	)

	roomHandler := handlers.NewRoomHandler(
		services.NewRoomService(repositories.NewRoomRepository(db)),
	)

	roomTypeHandler := handlers.NewRoomTypeHandler(
		services.NewRoomTypeService(repositories.NewRoomTypeRepository(db)),
	)

	guestHandler := handlers.NewGuestHandler(
		services.NewGuestService(repositories.NewGuestRepository(db)),
	)

	paymentHandler := handlers.NewPaymentHandler(
		services.NewPaymentService(repositories.NewPaymentRepository(db)),
	)

	bookingHandler := handlers.NewBookingHandler(
		services.NewBookingService(repositories.NewBookingRepository(db)),
	)

	staffHandler := handlers.NewStaffHandler(
		services.NewStaffService(
			repositories.NewStaffRepository(
				db,
				repositories.NewHotelRepository(db),
				repositories.NewUserRepository(db),
			),
		),
	)

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
	}

	publicRoutes := []struct {
		route   string
		handler interface{}
	}{
		{"/api/hotels", hotelHandler},
		{"/api/rooms", roomHandler},
		{"/api/room-types", roomTypeHandler},
		{"/api/guests", guestHandler},
		{"/api/payments", paymentHandler},
		{"/api/bookings", bookingHandler},
	}

	for _, r := range publicRoutes {
		group := router.Group(r.route)
		switch h := r.handler.(type) {
		case *handlers.HotelHandler:
			group.GET("", h.GetAllHotels)
			group.GET(":id", h.GetHotelByID)
		case *handlers.RoomHandler:
			group.GET("", h.GetAllRooms)
			group.GET(":id", h.GetRoomByID)
		case *handlers.RoomTypeHandler:
			group.GET("", h.GetAllRoomTypes)
			group.GET(":id", h.GetRoomTypeByID)
		case *handlers.GuestHandler:
			group.GET("", h.GetAllGuests)
			group.GET(":id", h.GetGuestByID)
		case *handlers.PaymentHandler:
			group.GET("", h.GetAllPayments)
			group.GET(":id", h.GetPaymentByID)
		case *handlers.BookingHandler:
			group.GET("", h.GetAllBookings)
			group.GET(":id", h.GetBookingByID)
		}
	}

	protected := router.Group("/api", middlewares.CheckJwt(), middlewares.IsAdmin())
	{
		protected.GET("/profile", authHandler.Profile)

		registerCRUDRoutes(protected.Group("/hotels"), hotelHandler)
		registerCRUDRoutes(protected.Group("/rooms"), roomHandler)
		registerCRUDRoutes(protected.Group("/room-types"), roomTypeHandler)
		registerCRUDRoutes(protected.Group("/guests"), guestHandler)
		registerCRUDRoutes(protected.Group("/payments"), paymentHandler)
		registerCRUDRoutes(protected.Group("/bookings"), bookingHandler)
	}

	staffRoutes := router.Group("/api/staffs", middlewares.CheckJwt(), middlewares.IsStaff())
	{
		staffRoutes.POST("", staffHandler.CreateStaff)
		staffRoutes.GET("", staffHandler.GetAllStaff)
		staffRoutes.GET(":id", staffHandler.GetStaffByID)
		staffRoutes.PUT(":id", staffHandler.UpdateStaff)
		staffRoutes.DELETE(":id", staffHandler.DeleteStaff)
	}

	return router
}

func registerCRUDRoutes(group *gin.RouterGroup, handler interface{}) {
	switch h := handler.(type) {
	case *handlers.HotelHandler:
		group.POST("", h.CreateHotel)
		group.PUT(":id", h.UpdateHotel)
		group.DELETE(":id", h.DeleteHotel)
	case *handlers.RoomHandler:
		group.POST("", h.CreateRoom)
		group.PUT(":id", h.UpdateRoom)
		group.DELETE(":id", h.DeleteRoom)
	case *handlers.RoomTypeHandler:
		group.POST("", h.CreateRoomType)
		group.PUT(":id", h.UpdateRoomType)
		group.DELETE(":id", h.DeleteRoomType)
	case *handlers.GuestHandler:
		group.POST("", h.CreateGuest)
		group.PUT(":id", h.UpdateGuest)
		group.DELETE(":id", h.DeleteGuest)
	case *handlers.PaymentHandler:
		group.POST("", h.CreatePayment)
		group.PUT(":id", h.UpdatePayment)
		group.DELETE(":id", h.DeletePayment)
	case *handlers.BookingHandler:
		group.POST("", h.CreateBooking)
		group.PUT(":id", h.UpdateBooking)
		group.DELETE(":id", h.DeleteBooking)
	}
}
