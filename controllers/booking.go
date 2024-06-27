package controllers

import (
	"net/http"

	"github.com/adityagoyal597/hotel-bookings/models"
	"github.com/gin-gonic/gin"
)

func CreateBooking(context *gin.Context) {
	var booking models.Booking
	err := context.ShouldBindJSON(&booking)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Data"})
		return
	}

	userID, exists := context.Get("userID")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	booking.UserID = userID.(int)

	err = booking.Save()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could Not Save Booking"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Booking Created Successfully!"})
}

func GetBookings(context *gin.Context) {
	userID, exists := context.Get("userID")

	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	bookings, err := models.GetBookingsByUserID(userID.(int))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Get Bookings"})
		return
	}

	context.JSON(http.StatusOK, bookings)
}
