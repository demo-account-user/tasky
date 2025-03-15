package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = CreateMongoClient()

func CreateMongoClient() *mongo.Client {
	godotenv.Overload()
	MongoDbURI := os.Getenv("MONGODB_URI")
	uname := os.Getenv("MONGODB_USERNAME")
	passwd := os.Getenv("MONGODB_PASSWORD")

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource: "admin",
		Username: uname,
		Password: passwd,
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDbURI).SetAuth(credential))
	if err != nil {
		log.Fatal(err)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	fmt.Println("Connected to MONGO -> ", MongoDbURI)
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("go-mongodb").Collection(collectionName)
}
