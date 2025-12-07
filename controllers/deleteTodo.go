package controllers

import (
	"errors"
	"fmt"
	"gorouter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required").Error()})
		return
	}

	for index, v := range models.Todos {
		if v.ID == id {
			models.Todos = append(models.Todos[:index], models.Todos[index+1:]...)
			c.JSON(http.StatusAccepted, id)
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("todo with id (%s) not found", id).Error()})

}
