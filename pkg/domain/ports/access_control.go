package ports

import "github.com/jairogloz/go-budget/pkg/domain/core"

// AccessControlService exposes methods to get the access control for a given user.
type AccessControlService interface {
	GetFeatureAccess(userLevel core.UserLevel) (core.FeatureAccess, error)

	// AuthenticateUser authenticates a user using OAuth 2.0 and returns user information.
	AuthenticateUser(provider string, code string) (core.UserInfo, error)
}
