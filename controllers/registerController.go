package controllers

import (
	"go-gin-crud/config"
	"go-gin-crud/helpers"
	"go-gin-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// validate PostForm
	err := helpers.ValidateForm(c, "name", "email", "username", "password")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	tx := config.DB.Begin()
	q := models.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Username: c.PostForm("username"),
		Password: helpers.HashPassword(c.PostForm("password")),
	}
	if err := config.DB.Save(&q).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Product",
		"data":    q,
	})
}
