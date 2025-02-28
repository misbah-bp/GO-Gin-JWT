package main

import (
	"github.com/gin-gonic/gin"
	"guthyb.com/misbah-bp/JWT-authentication/controllers"
	"guthyb.com/misbah-bp/JWT-authentication/initilaizers"
	"guthyb.com/misbah-bp/JWT-authentication/middleware"
)

func init() {
	initilaizers.LoadEnvVariables()
	initilaizers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/singUp", controllers.SingupUser)
	r.POST("/singin", controllers.SingInUser)
	r.GET("/validate", middleware.RequiredAUth, controllers.Validate)
	r.POST("/logout", controllers.Logout)
	r.Run()
}
