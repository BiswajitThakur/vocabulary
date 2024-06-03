package router

import (
	"context"
	"log"

	"github.com/BiswajitThakur/vocabulary/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// type Voca struct {
// 	Word     string   `form:"word" json:"word"`
// 	Type     string   `form:"type" json:"type"`
// 	Meaning  string   `form:"meaning" json:"meaning"`
// 	Synonym  []string `form:"synonym[]" json:"synonym"`
// 	Antonym  []string `form:"antonym[]" json:"antonym"`
// 	Examples []string `form:"examples[]" json:"examples"`
// }

func SetData(c *gin.Context, collection *mongo.Collection) {
	var myw db.Voca
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&myw) == nil {
		myw.ID = primitive.NewObjectID()
		insertResult, err := collection.InsertOne(context.TODO(), myw)
		log.Println("====== Bind By Query String ======")
		log.Println(myw.Word)
		log.Println(myw.Type)
		log.Println(myw.Meaning)
		log.Println(myw.Synonym)
		log.Println(myw.Antonym)
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
