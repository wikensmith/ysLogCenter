package ysHTTP

import (
	"fmt"
	"log"
	"logCenter/ysMongo"

	"github.com/gin-gonic/gin"
)

func handlerAdd(context *gin.Context) {
	log.Println("hello gin start")
	//fmt.Println(ioutil.ReadAll(context.Request.Body))
	if err := ysMongo.AddToMongo(context); err != nil {
		fmt.Println("error in handler", err)
	}

	context.JSON(200, gin.H{
		"code":    200,
		"success": true,
		"msg":     "test msg here",
	})
}

func handlerQuery(context *gin.Context) {
	fmt.Println("in handlerQuery", context.Request.Body)
}

func HttpMain() {
	router := gin.Default()
	fmt.Println("test begin here")
	router.POST("/add", handlerAdd)
	router.POST("/query", handlerQuery)
	b := router.Run("localhost:9090")
	fmt.Println(b)
}
