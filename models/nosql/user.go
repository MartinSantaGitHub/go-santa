package nosql

import "go.mongodb.org/mongo-driver/bson/primitive"

/* User model for the mongo DB */
type User struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Active bool               `bson:"active"`
}
