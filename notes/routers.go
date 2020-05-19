package main

import (
	"github.com/gin-gonic/gin"
	"notes/controllers"
	"notes/middleware"
)

func Routers(r *gin.Engine) *gin.Engine{
	r.Use(middleware.Cors(),middleware.RecoverMiddleware())


	userController := controllers.NewUser()
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	return r
}
