package access_control

import (
	"fmt"
	"github.com/jairogloz/go-budget/pkg/domain/core"
	"github.com/jairogloz/go-budget/pkg/domain/ports"
)

var _ ports.AccessControlService = (*Service)(nil)

// Service implements ports.AccessControlService.
type Service struct {
}

// NewService creates a new access control service.
func NewService() *Service {
	return &Service{}
}

// GetFeatureAccess returns the feature access for a given user level.
func (s Service) GetFeatureAccess(userLevel core.UserLevel) (core.FeatureAccess, error) {
	// FeatureAccessMapDefault is the default feature access map.
	var FeatureAccessMapDefault = core.FeatureAccessMap{
		core.UserLevelFree:  {MaxAccounts: 3, MaxCategories: 3},
		core.UserLevelPaid1: {MaxAccounts: 0, MaxCategories: 0},
	}

	if featureAccess, ok := FeatureAccessMapDefault[core.UserLevel(userLevel)]; ok {
		return featureAccess, nil
	}

	return core.FeatureAccess{}, fmt.Errorf("invalid user level %s", userLevel)
}

// AuthenticateUser authenticates a user using OAuth 2.0 and returns user information.
// Todo: implement the OAuth 2.0 authentication.
func (s Service) AuthenticateUser(provider string, code string) (core.UserInfo, error) {
	return core.UserInfo{
		ID: "123456",
	}, nil
}
