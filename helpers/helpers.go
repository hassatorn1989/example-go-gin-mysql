package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidateForm and ValidateFormParams are helper functions to validate form data
func ValidateForm(c *gin.Context, fields ...string) error {
	for _, field := range fields {
		if c.PostForm(field) == "" {
			return fmt.Errorf("%s is required", field)
		}
	}
	return nil
}

func ValidateFormParams(c *gin.Context, fields ...string) error {
	for _, field := range fields {
		if c.Param(field) == "" {
			return fmt.Errorf("%s is required", field)
		}
	}
	return nil
}
