package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
	"net/http"
	"strings"
)

var supabaseJwtSecret = []byte("<jwt-secret>")

// AuthRequired for now, simply inserts a user ID into the context.
func (h *Handler) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return supabaseJwtSecret, nil
		})

		// If token is invalid or parsing fails
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: subject claim not found"})
			c.Abort()
			return
		}

		user := &pkgCore.User{
			ID:    sub,
			Level: pkgCore.UserLevelFree,
			//UserInfo: &userInfo,
		}

		featureAccess, err := h.AccessControlService.GetFeatureAccess(user.Level)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		user.FeatureAccess = &featureAccess

		// Todo: Validate user has access to path

		// Inject the user into the context
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), pkgCore.CtxKeyUser, user))

		c.Next()
	}
}
