package main

import (
	"Rental-car/internal/routes"
	"Rental-car/pkg/database"
	"Rental-car/pkg/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	db := database.ConnectToMysql()

	if db != nil {
		fmt.Println("Connected")
	}

	r.GET("/ping", middleware.AuthJWT(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Routes
	routes.UserRoutes(r, db)
	routes.CarsRoutes(r, db)
	routes.BookingRoutes(r, db)

	r.Run() // listen and serve on 0.0.0.0:8080
}
