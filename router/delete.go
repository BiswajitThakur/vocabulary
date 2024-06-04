package router

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteById(c *gin.Context, collection *mongo.Collection) {
	var idStr = c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if deleteResult.DeletedCount == 0 {
		c.JSON(400, gin.H{"_id": idStr, "status": "Not Found"})
		return
	}
	c.JSON(400, gin.H{"_id": idStr, "status": "Deleted"})
}
