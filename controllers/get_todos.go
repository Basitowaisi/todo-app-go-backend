package controllers

import (
	"gorouter/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {

	c.JSON(http.StatusOK, models.Todos)
}
