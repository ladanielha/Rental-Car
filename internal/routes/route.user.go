package routes

import (
	"Rental-car/internal/handlers"
	"Rental-car/internal/repositories"
	"Rental-car/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repositories.NewRepoUser(db)
	serv := services.NewServiceUser(repo)
	handler := handlers.NewHandlerUser(serv)

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
}
