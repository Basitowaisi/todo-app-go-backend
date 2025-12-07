package controllers

import (
	"gorouter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {

	response := []models.Todo{}

	for _, v := range models.Todos {
		response = append(response, v)
	}

	c.JSON(http.StatusOK, response)
}
