package router

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllWords(c *gin.Context, collection *mongo.Collection) {
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(500, gin.H{"status": "Somthing Went Wrong"})
		log.Println(err)
		return
	}
	var collection_words []bson.M
	if err = cursor.All(context.TODO(), &collection_words); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, collection_words)
}
