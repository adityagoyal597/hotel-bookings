package controllers

import (
	"net/http"
	"strconv"

	"github.com/adityagoyal597/hotel-bookings/models"
	"github.com/gin-gonic/gin"
)

func GetRooms(context *gin.Context) {
	rooms, err := models.GetAllRooms()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get rooms"})
		return
	}

	context.JSON(http.StatusOK, rooms)
}

func CreateRoom(context *gin.Context) {
	var room models.Room
	err := context.ShouldBindJSON(&room) // ASKING DATA FROM THE USER

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Data"})
		return
	}

	err = room.Save() //METHOD

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Save Room"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Room Created Successfully!"})
}

func UpdateRoom(context *gin.Context) {
	var room models.Room // SHOULD DECLARE WHILE UPDATING , CREATING
	roomId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Room ID"})
		return
	}

	err = context.ShouldBindJSON(&room)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Data"})
		return
	}

	room.ID = roomId

	err = room.Update() // METHOD

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update room"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Room Updated Successfully!"})
}

func DeleteRoom(context *gin.Context) {
	roomId, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Room ID"})
		return
	}

	err = models.Delete(roomId) // FUNCTION

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete room"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Room Deleted Successfully!"})
}
