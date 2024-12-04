package routes

import (
	"Rental-car/internal/handlers"
	"Rental-car/internal/repositories"
	"Rental-car/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookingRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRepoBooking(db)
	repoCar := repositories.NewRepoCars(db)
	serv := services.NewServiceBooking(repo, repoCar)
	handler := handlers.NewHandlerBooking(serv)

	bookingRoute := r.Group("/bookings")
	{
		bookingRoute.POST("", handler.CreateBooking)
		bookingRoute.GET("", handler.GetBookings)
		bookingRoute.GET(":id", handler.GetBooking)
		bookingRoute.PUT(":id", handler.UpdateBooking)
	}
}
