package middleware

import (
	"net/http"

	"example.com/Hotel_Bookings/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(context *gin.Context) {
	authHeader := context.Request.Header.Get("Authorization")

	if authHeader == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization Header is Required"})
		return
	}

	userID, err := utils.VerifyToken(authHeader)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, "Not Authorized")
		return
	}

	context.Set("userID", userID)

	context.Next()
}
