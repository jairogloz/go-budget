package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-budget/cmd/gin/core"
	"net/http"
	"os"
)

// AuthRequired for now, simply inserts a user ID into the context.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Here you would normally check for a valid token in the request headers and validate it
		// For simplicity, let's just simulate an authenticated user with a boolean
		isAuthenticated := true

		if !isAuthenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Inject the user ID into the context
		// Todo: change this to change userID
		testUserID := "1"
		if os.Getenv("SERVER_MODE") == "development" {
			testUserID = "test"
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), core.UserIDKey, testUserID))

		c.Next()
	}
}
