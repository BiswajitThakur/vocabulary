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
	//dbs.InsertOne(context.TODO(), )
	router := gin.Default()
	//router.Static("/", "./public")
	router.POST("/set_data", func(ctx *gin.Context) {
		r.SetData(ctx, dbs)
	})
	// router.POST("/get_word", func(ctx *gin.Context) {
	// 	r.GetWordById(ctx, dbs)
	// })
	router.GET("/words", func(ctx *gin.Context) {
		r.GetAllWords(ctx, dbs)
	})
	router.GET("/words/:id", func(ctx *gin.Context) {
		r.GetWordById(ctx, dbs)
	})

	router.Static("/src", "./public")
	router.Run() // listen and serve on 0.0.0.0:8080
	// // db, _ := c.Open("clover-db")
	// //store, _ := badgerstore.Open(badger.DefaultOptions("").WithInMemory(true))
	// //db, _ := c.OpenWithStore(store)
	// // db, err := badger.Open(badger.DefaultOptions("./tmp-db"))
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // defer db.Close()
}
