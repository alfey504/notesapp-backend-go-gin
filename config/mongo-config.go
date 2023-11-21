package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Client {
	fmt.Println("Connecting to mongo servers ...")

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	mongoUri := GetEnvVariableFatal(MONGODB_URI)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	/*
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()
	*/

	var result bson.M

	if err := client.Database("admin").RunCommand(context.TODO(), bson.M{"ping": 1}).Decode(&result); err != nil {
		panic(err)
	}

	fmt.Println("Connected to mongoDb")
	return client
}

var DB *mongo.Client = ConnectMongo()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := GetEnvVariableFatal(DB_NAME)
	return client.Database(dbName).Collection(collectionName)
}

func MongoDisconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
