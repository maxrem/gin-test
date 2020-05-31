package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func getLogin(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"login.html",
		gin.H {
			"title": "Login",
		},
	)
}

func postLogin(context *gin.Context) {
	var form Login

	if err := context.ShouldBind(&form); err != nil {
		context.HTML(http.StatusBadRequest, "login.html", gin.H{"status": err.Error()})

		return
	}

	if form.Email != "test@example.com" || form.Password != "pass" {
		context.HTML(http.StatusUnauthorized, "login.html", gin.H{"status": "unauthorized"})

		return
	}

	context.HTML(http.StatusOK, "login.html", gin.H{"status": "logged in!"})
}