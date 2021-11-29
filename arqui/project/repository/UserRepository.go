package repository

import (
	d "arqui/project/db"
	u "arqui/project/model"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//creates user
func CreateUser(user *u.User) (primitive.ObjectID, error) {
	//var newUser User
	client, ctx, cancel := d.GetConnection()
	defer cancel()

	defer cancel()
	defer client.Disconnect(ctx)
	user.ID = primitive.NewObjectID()

	result, err := client.Database("auth_db").Collection("user").InsertOne(ctx, user)

	if err != nil {
		log.Printf("Could not create user: %v", err)
		return primitive.NilObjectID, err
	}
	//object id
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func FindUsers() ([]*u.User, error) {
	var users []*u.User

	client, ctx, cancel := d.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("auth_db")
	collection := db.Collection("user")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}

	return users, nil
}

func FindUserByNumber(userNumber string) (*u.User, error) {
	var user *u.User

	client, ctx, cancel := d.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("auth_db")
	collection := db.Collection("user")
	filter := bson.D{{Key: "number", Value: userNumber}}

	result := collection.FindOne(ctx, filter)
	if result == nil {
		return nil, errors.New("Could not find a user")
	}

	err := result.Decode(&user)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Tasks: %v", user)
	return user, nil
}
