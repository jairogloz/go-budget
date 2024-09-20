package core

// Category reflects a category in the system.
type Category struct {
	ID     string `json:"id" bson:"_id"`
	UserId string `json:"user_id" bson:"user_id"`
}
