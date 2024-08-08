package controllers

import (
	"go-gin-crud/config"
	"go-gin-crud/helpers"
	"go-gin-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User struct
func Index(c *gin.Context) {
	var q []models.User
	config.DB.Find(&q)
	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Products",
		"data":    q,
	})
}

func Create(c *gin.Context) {
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

func Show(c *gin.Context) {
	var q models.User
	config.DB.First(&q, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Product",
		"data":    q,
	})
}

func Update(c *gin.Context) {
	err := helpers.ValidateForm(c, "name", "email", "username")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = helpers.ValidateFormParams(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tx := config.DB.Begin()
	var q models.User
	config.DB.First(&q, c.Param("id"))
	q.Name = c.PostForm("name")
	q.Email = c.PostForm("email")
	q.Username = c.PostForm("username")
	if c.PostForm("password") != "" {
		q.Password = helpers.HashPassword(c.PostForm("password"))
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
		"message": "Update Product",
		"data":    q,
	})
}

func Destroy(c *gin.Context) {
	err := helpers.ValidateFormParams(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tx := config.DB.Begin()
	var q models.User
	config.DB.First(&q, c.Param("id"))
	if err := config.DB.Delete(&q).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Product",
		"data":    q,
	})
}
