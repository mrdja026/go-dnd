package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("URL:MONGO_DB_CONNECTION"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx
}

func GetData() {
	client, ctx := ConnectToMongo()

	quickstartDatabase := client.Database("sample_mflix")
	moviesCollection := quickstartDatabase.Collection("movies")

	cursor, err := moviesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Find error", err)
	}
	var movies []bson.M
	if err = cursor.All(ctx, &movies); err != nil {
		log.Fatal("Cursor all", err)
	}
	defer client.Disconnect(ctx)

	fmt.Println(movies)
}
