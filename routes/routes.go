package routes

import (
	"go-gin-crud/controllers"
	"go-gin-crud/middlewares"

	"github.com/gin-gonic/gin"
)

// Route struct
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)

		v1.POST("/logout", controllers.Logout)

		protectedRoutes := v1.Group("/protected")
		protectedRoutes.Use(middlewares.JWTAuthMiddleware())
		{
			user := protectedRoutes.Group("/users")
			{
				user.GET("/", controllers.Index)
				user.POST("/", controllers.Create)
				user.GET("/:id", controllers.Show)
				user.PUT("/:id", controllers.Update)
				user.DELETE("/:id", controllers.Destroy)
			}
		}
	}
	return r
}
