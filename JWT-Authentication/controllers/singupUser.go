package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"guthyb.com/misbah-bp/JWT-authentication/initilaizers"
	"guthyb.com/misbah-bp/JWT-authentication/models"
)

func SingupUser(c *gin.Context) {
	// get the data from the body
	var body struct {
		Email    string
		Password string
		Name     string
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// hash passwrd
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 5)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "problem in hasing "})
		return
	}

	// create user
	user := models.User{Email: body.Email, Password: string(hash), Name: body.Name}

	result := initilaizers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not created ",
		})
	} else {

		// responce
		c.JSON(http.StatusOK, gin.H{
			"message": "user is created",
			"User":    user,
		})
	}

}
