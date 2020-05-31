package main

import "github.com/gin-gonic/gin"

func main() {
	ginEngine := gin.Default()

	ginEngine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginEngine.Run()
}
