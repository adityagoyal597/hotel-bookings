package main

import (
	"example.com/Hotel_Bookings/db"
	"example.com/Hotel_Bookings/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
