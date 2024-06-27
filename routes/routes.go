package routes

import (
	"example.com/Hotel_Bookings/controllers"
	"example.com/Hotel_Bookings/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", controllers.Signup)

	server.POST("/login", controllers.Login)

	authorized := server.Group("/")

	authorized.Use(middleware.AuthMiddleware)

	authorized.GET("/rooms", controllers.GetRooms)

	authorized.POST("rooms", controllers.CreateRoom)

	authorized.PUT("/rooms/:id", controllers.UpdateRoom)

	authorized.DELETE("/rooms/:id", controllers.DeleteRoom)

	authorized.POST("/bookings", controllers.CreateBooking)

	authorized.GET("/bookings", controllers.GetBookings)
}
