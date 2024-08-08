package controllers

import (
	"fmt"
	"go-gin-crud/config"
	"go-gin-crud/helpers"
	"go-gin-crud/middlewares"
	"go-gin-crud/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// validate PostForm
	err := helpers.ValidateForm(c, "username", "password")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// check user
	var q models.User
	config.DB.Where("username = ?", c.PostForm("username")).First(&q)
	if q.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}
	// check password
	if !helpers.CheckPasswordHash(c.PostForm("password"), q.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Password",
		})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": q.Username,
		"email":    q.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"token":   tokenString,
	})
}

func Logout(c *gin.Context) {
	tokenString := strings.Split(c.GetHeader("Authorization"), " ")[1]
	fmt.Println(tokenString)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is requiredcccc"})
		return
	}

	middlewares.AddToBlacklist(tokenString)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Success",
	})

}
