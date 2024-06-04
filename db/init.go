package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Voca struct {
	ID       primitive.ObjectID `form:"id" json:"_id" bson:"_id"`
	Word     string             `form:"word" json:"word" bson:"word"`
	Type     string             `form:"type" json:"type" bson:"type"`
	Meaning  string             `form:"meaning" json:"meaning" bson:"meaning"`
	Synonym  []string           `form:"synonym[]" json:"synonym" bson:"synonym"`
	Antonym  []string           `form:"antonym[]" json:"antonym" bson:"antonym"`
	Examples []string           `form:"examples[]" json:"examples" bson:"examples"`
}

func Init() (*mongo.Collection, *mongo.Collection) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	database := os.Getenv("DATABASE")
	if database == "" {
	    log.Fatal("DATABASE: environment variable not found.")
	}
	collection := os.Getenv("COLLECTION")
	if collection == "" {
	    log.Fatal("COLLECTION: environment variable not found.")
	}
	admins := client.Database(database).Collection("admin")
	adminsCollection := client.Database(database).Collection(collection)
	return admins, adminsCollection
}
