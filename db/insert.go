package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(collection *mongo.Collection, data primitive.D) interface{} {
	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		panic(err)
	}
	return result.InsertedID
}
