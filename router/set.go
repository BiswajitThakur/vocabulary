package router

import (
	"context"
	"log"

	"github.com/BiswajitThakur/vocabulary/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetData(c *gin.Context, collection *mongo.Collection) {
	var myw db.Voca
	if c.ShouldBind(&myw) == nil {
		myw.ID = primitive.NewObjectID()
		insertResult, err := collection.InsertOne(context.TODO(), myw)
		if err != nil {
			c.JSON(200, gin.H{"status": "faild"})
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{"status": "success", "id": insertResult.InsertedID})
		return
	}

	c.JSON(200, gin.H{"status": "invalid request"})
}
