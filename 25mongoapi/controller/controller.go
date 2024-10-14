package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const dbName = "netflix"

const colName = "watchlist"

//most important

var collection *mongo.Collection

//connect with db

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection Success")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance is ready")
}
