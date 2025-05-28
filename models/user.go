package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Age    int                `json:"age" bson:"age"`
	Gender string             `json:"gender" bson:"gender"`
}
