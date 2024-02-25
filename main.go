package main

import (
	"go-jwt-auth-api/controllers"
	"go-jwt-auth-api/initializers"
	"go-jwt-auth-api/middleware"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.OpenDb()

	router := gin.Default()

	router.GET("/", controllers.Home)
	router.GET("/private", middleware.RequireAuth, controllers.Private)
	router.GET("/logout", controllers.Logout)
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)

	router.Run(":" + os.Getenv("PORT"))
}
