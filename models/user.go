package models

import "go-gin-crud/config"

// Product struct
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// TableName will allow us to define a custom table name
func (b *User) TableName() string {
	return "users"
}

func Migrate() {
	config.DB.AutoMigrate(&User{})
}
