package core

// FeatureAccess represents the access to features for a user level.
// For most integer fields, 0 means unlimited.
type FeatureAccess struct {
	MaxAccounts   int `json:"max_accounts" bson:"max_accounts"`
	MaxCategories int `json:"max_categories" bson:"max_categories"`
}

// FeatureAccessMap represents the access to features for all user levels.
type FeatureAccessMap map[UserLevel]FeatureAccess
