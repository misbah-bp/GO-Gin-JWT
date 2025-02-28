package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"guthyb.com/misbah-bp/JWT-authentication/initilaizers"
	"guthyb.com/misbah-bp/JWT-authentication/models"
)

func SingInUser(c *gin.Context) {
	// resived boday data
	var body struct {
		Email    string
		Password string
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// find user

	fmt.Println("user email: ", body.Email)

	// If email is empty, return error
	if body.Email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}

	// check the user in db
	var user models.User
	fmt.Println("my users user: ", user)

	initilaizers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password or email"})
		return
	}

	// compare th hash password with paswod
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password "})
		return
	}

	// genrate JWt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRATE")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "problem to genarte JWT token "})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}
