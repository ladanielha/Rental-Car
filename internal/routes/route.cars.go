package routes

import (
	"Rental-car/internal/handlers"
	"Rental-car/internal/repositories"
	"Rental-car/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CarsRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRepoCars(db)
	serv := services.NewServiceCars(repo)
	handler := handlers.NewHandlerCars(serv)

	carsRoute := r.Group("/cars")
	{
		carsRoute.POST("", handler.CreateCar)
		carsRoute.PUT(":id")
	}
}
