package main

import "github.com/gin-gonic/gin"

func main() {
	run := gin.Default()
	run.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"hello": "Hello World",
		})
	})
	_ = run.Run(":8181")
}
