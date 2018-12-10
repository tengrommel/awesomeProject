package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)

	clientOptions := options.ClientOptions{}
	if client, err = mongo.Connect(context.TODO(), "mongodb://127.0.0.1:27017", clientOptions.SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Println(err)
		return
	}
	database = client.Database("my_db")
	collection = database.Collection("my_collection")
	collection = collection
}
