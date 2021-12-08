package model

//import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     string  `bson:"id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Number string             `bson:"number,omitempty"`
}
