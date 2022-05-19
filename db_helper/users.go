package db_helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserCollection() mongo.Collection {

	mongo_url := "mongodb://" + os.Getenv("MONGO_URL") + ":27017/"

	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return *client.Database("did_homa").Collection("users_collection")
}

type User struct {
	UserId   int64
	Warnings int8
	Checked  bool
}

func NewUser(id int64) User {

	users_collection := UserCollection()

	var u User

	if !u.IsExists(id) {
		insertResult, err := users_collection.InsertOne(context.TODO(), User{id, 0, false})
		log.Println("Inserted a single document: ", insertResult.InsertedID)
		if err != nil {
			log.Fatal(err)

		}
	}

	var user User
	users_collection.FindOne(context.TODO(), bson.D{{"userid", id}}).Decode(&user)
	return user
}

func (u *User) IsExists(userId int64) bool {

	users_collection := UserCollection()

	var user User
	users_collection.FindOne(context.TODO(), bson.D{{"userid", userId}}).Decode(&user)
	return user.UserId != 0

}

func AddWarning(userId int64) {

	users_collection := UserCollection()

	filter := bson.D{{"userid", userId}}

	var user User
	users_collection.FindOne(context.TODO(), filter).Decode(&user)

	//Update one document
	update := bson.D{
		{"$set", bson.D{
			{"warnings", user.Warnings + 1},
		}},
	}

	updateResult, err := users_collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func AddChecked(userId int64) {

	users_collection := UserCollection()

	filter := bson.D{{"userid", userId}}

	var user User
	users_collection.FindOne(context.TODO(), filter).Decode(&user)

	//Update one document
	update := bson.D{
		{"$set", bson.D{
			{"checked", true},
		}},
	}

	updateResult, err := users_collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
