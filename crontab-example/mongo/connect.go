package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	cli, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := cli.Database("testing").Collection("numbers")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	fmt.Println("Id: ", id)
	//ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	//err = cli.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	panic(err)
	//}
}
