package controllers

import (
	"errors"
	"fmt"
	"gorouter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required").Error()})
		return
	}

	for index, value := range models.Todos {
		if value.ID == id {
			models.Todos[index] = models.Todo{
				ID:          value.ID,
				Todo:        value.Todo,
				IsCompleted: true,
			}

			c.JSON(http.StatusAccepted, models.Todos[index])
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("todo with id (%s) not found", id).Error()})

}
