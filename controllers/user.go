package controllers

import (
	"net/http"

	"example.com/Hotel_Bookings/models"
	"example.com/Hotel_Bookings/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) { // CREATING A USER
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Save User"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully!"})
}

func Login(context *gin.Context) { // LOGIN OF A USER TO RETURN A TOKEN
	var loginUser models.User // loginUser STRUCT

	err := context.ShouldBindJSON(&loginUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request Data"})
		return
	}

	user, err := models.GetUserByEmail(loginUser.Email) // user STRUCT THAT GETS RETURN BY GetUserByID FUNCTION

	if user == nil || err != nil { // CHECKING WHETHER THE USER IS IN THE DATABASE OF USER
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		return
	}

	err = user.CheckPassword(loginUser.Password) //METHOD
	// CHECKS WHEHTER THE PASSWORD ENTERED BY THE loginUser IS SAME AS THE user Password

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Generate Token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successfull!", "token": token})

}
