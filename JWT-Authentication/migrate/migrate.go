package main

import (
	"guthyb.com/misbah-bp/JWT-authentication/initilaizers"
	"guthyb.com/misbah-bp/JWT-authentication/models"
)

func init() {
	initilaizers.LoadEnvVariables()
	initilaizers.ConnectToDb()
}

func main() {

	initilaizers.DB.AutoMigrate(&models.User{})
}
