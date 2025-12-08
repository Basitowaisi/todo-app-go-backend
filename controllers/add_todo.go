package controllers

import (
	"errors"
	"fmt"
	"gorouter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var item models.Todo
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := models.ValidateDuplicateInput(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Executing after validation....")

	stId := models.GenerateId()

	if item.Todo == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("you need to provide us with a todo").Error()})
		return
	}

	newTodo := models.Todo{
		ID:          stId,
		Todo:        item.Todo,
		IsCompleted: item.IsCompleted || false,
	}

	models.Todos = append(models.Todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}
