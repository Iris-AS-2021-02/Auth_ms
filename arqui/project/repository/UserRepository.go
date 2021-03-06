package repository

import (
	d "arqui/project/db"
	u "arqui/project/model"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
	uri                      = "mongodb+srv://root:2021@cluster0.34iec.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	uri2                     = "mongodb+srv://root:2021@cluster0.34iec.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

//creates user
func CreateUser(user *u.User) (primitive.ObjectID, error) {
	//var newUser User
	client, ctx, cancel := d.GetConnection(uri)
	defer cancel()
	defer client.Disconnect(ctx)
	user.ID = primitive.NewObjectID().Hex()
	result, err := client.Database("auth_db").Collection("user").InsertOne(ctx, user)
	if err != nil {
		log.Printf("Could not create user: %v", err)
		return primitive.NilObjectID, err
	}

	//object id
	oid := result.InsertedID.(primitive.ObjectID)

	client2, ctx2, cancel2 := d.GetConnection(uri2)
	defer cancel2()
	defer client2.Disconnect(ctx)
	result2, err2 := client2.Database("auth_db_2").Collection("user").InsertOne(ctx2, user)
	if err2 != nil {
		log.Printf("Could not create user: %v", err)
		return primitive.NilObjectID, err
	}

	//object id
	oid2 := result2.InsertedID.(primitive.ObjectID)

	log.Printf("Could create user: %v", oid2)
	return oid, nil
}

func FindUsers() ([]*u.User, error) {
	var users []*u.User

	client, ctx, cancel := d.GetConnection(uri)
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database("auth_db")
	db2 := client.Database("auth_db_2")
	collection := db.Collection("user")
	collection2 := db2.Collection("user")
	cursor, err := collection.Find(ctx, bson.D{})
	cursor2, err2 := collection2.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	defer cursor.Close(ctx)
	defer cursor2.Close(ctx)
	err = cursor.All(ctx, &users)
	err2 = cursor2.All(ctx, &users)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	if err2 != nil {
		log.Printf("Failed marshalling %v", err2)
		return nil, err
	}
	return users, nil
}

func FindUserByNumber(userNumber string) (*u.User, error) {
	var user *u.User

	client, ctx, cancel := d.GetConnection(uri)
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

func FindUsersWithNumber(numbers string) ([]*u.User, error) {
	users, err := FindUsers()
	if err != nil {
		return nil, err
	}
	inNumbers := map[string]int{
		"dummy": 1,
	}
	i := 0
	cur := ""
	for i < len(numbers) {
		if numbers[i] == ',' {
			inNumbers[cur] = 1
			cur = ""
		} else {
			cur += string(numbers[i])
		}
		i += 1
	}

	if cur != "" {
		inNumbers[cur] = 1
	}

	var valUsers []*u.User
	i = 0
	for i < len(users) {
		if inNumbers[users[i].Number] != 0 {
			valUsers = append(valUsers, users[i])
		}
		i += 1
	}
	return valUsers, nil
}
