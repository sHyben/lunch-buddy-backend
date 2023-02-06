package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sHyben/lunch-buddy-backend/pkg/lunch-buddy-backend/crypto"
	"net/http"
)

// AuthRequired is a middleware that checks if the request has a valid token
// It is called by router.Setup
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}
