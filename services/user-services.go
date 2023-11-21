package services

import (
	"context"
	"notes-appapi/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = config.GetCollection(config.DB, "users")

func UserExists(ctx context.Context, username string) (bool, error) {
	filter := bson.M{"username": username}
	user := userCollection.FindOne(ctx, filter)
	if user.Err() == mongo.ErrNoDocuments {
		println(user.Err().Error())
		return false, nil
	} else if user.Err() != nil {
		return true, user.Err()
	}
	return true, nil
}
