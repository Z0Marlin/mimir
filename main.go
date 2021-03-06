package main

import (
	"github.com/Z0marlin/mimir/controllers"
	"github.com/Z0marlin/mimir/db"
	"github.com/Z0marlin/mimir/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.Init()
	runMigrations()
	var userController controllers.UserController
	r.GET("/users/:id", userController.GET)
	r.GET("/users", userController.LIST)
	r.POST("/users", userController.POST)
	r.Run("localhost:8080")
}

func runMigrations() {
	db.D().AutoMigrate(&models.User{})
}
