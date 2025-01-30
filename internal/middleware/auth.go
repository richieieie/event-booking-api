package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richieieie/event-booking/internal/utils"
)

func Authenticate(c *gin.Context) {
	authorizedString := c.Request.Header.Get("Authorization")
	if authorizedString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Please login to create an event"})
		return
	}

	// Bearer token
	tokenType, tokenString := authorizedString[0:6], authorizedString[7:]
	if tokenType != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Please login to create an event"})
		return
	}

	userId, err := utils.VerifyJwtToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
