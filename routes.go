package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/", showIndexPage)

	userRoutes := router.Group("/u") // handlers.user.go
	{
		userRoutes.GET("/login", getLogin)
		userRoutes.POST("/login", postLogin)
	}
}

func showIndexPage(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}
