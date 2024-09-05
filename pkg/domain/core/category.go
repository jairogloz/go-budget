package core

// Category reflects a category in the system.
type Category struct {
	ID     interface{} `json:"id" bson:"_id"`
	Name   string      `json:"name" bson:"name"`
	UserId string      `json:"user_id" bson:"user_id"`
}
