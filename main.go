package main

import (
	_ "log"

	// badger "github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"

	// c "github.com/ostafen/clover"
	// badgerstore "github.com/ostafen/clover/v2/store/badger"
	db "github.com/BiswajitThakur/vocabulary/db"
	r "github.com/BiswajitThakur/vocabulary/router"
)

func main() {
	_, dbs := db.Init()
	router := gin.Default()
	router.POST("/set_data", func(ctx *gin.Context) {
		r.SetData(ctx, dbs)
	})
	router.GET("/words", func(ctx *gin.Context) {
		r.GetAllWords(ctx, dbs)
	})
	router.GET("/words/id/:id", func(ctx *gin.Context) {
		r.GetWordById(ctx, dbs)
	})
	router.DELETE("/words/id/:id", func(ctx *gin.Context) {
		r.DeleteById(ctx, dbs)
	})
	router.StaticFile("/", "./public/index.html")

	router.Static("/src", "./public")
	router.Run(":8080")

}
