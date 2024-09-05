package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	pkgCore "github.com/jairogloz/go-budget/pkg/domain/core"
)

// AuthRequired for now, simply inserts a user ID into the context.
func (h *Handler) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo: potentially move to an auth service
		// todo: implement actual authentication
		userInfo, err := h.AccessControlService.AuthenticateUser("google", "test")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// todo: query user from database
		user := &pkgCore.User{
			ID:       "1",
			Level:    pkgCore.UserLevelFree,
			UserInfo: &userInfo,
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
