package main

import (
	"gorouter/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", controllers.GetTodos)
	api.POST("/", controllers.AddTodo)
	api.PATCH("/:id", controllers.UpdateTodo)
	api.DELETE("/:id", controllers.DeleteTodo)

	router.Run(":4000")
}
