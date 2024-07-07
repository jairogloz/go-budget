package core

type UserLevel string

const (
	UserLevelFree  UserLevel = "free"
	UserLevelPaid1 UserLevel = "paid_1"
)

// User reflects a user in the system.
type User struct {
	ID    string    `json:"id" bson:"_id"`
	Level UserLevel `json:"level" bson:"level"`

	// FeatureAccess will be populated when the user is authenticated
	// based on the user level.
	FeatureAccess *FeatureAccess `json:"-" bson:"-"`

	// UserInfo will be populated when the user is authenticated.
	UserInfo *UserInfo `json:"-" bson:"-"`
}

// UserInfo represents the information returned by auth providers on
// successful authentication.
type UserInfo struct {
	// ID is the unique identifier of the user in the provider.
	ID string `json:"id"`
}
