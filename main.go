package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          string `json:"id" form:"id"`
	Todo        string `json:"todo" form:"todo"`
	IsCompleted bool   `json:"completed" form:"completed"`
}

// temporary data source
var todos = map[string]Todo{
	"1": {ID: "1", Todo: "Read the Book", IsCompleted: false},
	"2": {ID: "2", Todo: "Go to Market", IsCompleted: true},
}

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", GetTodos)
	api.POST("/", AddTodo)
	api.PATCH("/:id", UpdateTodo)
	api.DELETE("/:id", DeleteTodo)

	router.Run(":4000")

}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, ok := todos[id]
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("todo with id (%s) not found", id).Error()})
		return
	}

	for key := range todos {
		if key == id {
			delete(todos, id)
			break
		}
	}

	c.JSON(http.StatusAccepted, id)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	_, ok := todos[id]
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Errorf("todo with id (%s) not found", id).Error()})
		return
	}

	for key, value := range todos {
		if key == id {
			todos[id] = Todo{
				ID:          value.ID,
				Todo:        value.Todo,
				IsCompleted: true,
			}
			break
		}
	}

	c.JSON(http.StatusAccepted, todos[id])
}

func GetTodos(c *gin.Context) {

	response := []Todo{}

	for _, v := range todos {
		response = append(response, v)
	}

	c.JSON(http.StatusOK, response)
}

func AddTodo(c *gin.Context) {
	var item Todo
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	if err := validateDuplicateInput(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Executing after validation....")

	stId := generateId()

	if item.Todo == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("you need to provide us with a todo").Error()})
		return
	}

	todos[stId] = Todo{
		ID:          stId,
		Todo:        item.Todo,
		IsCompleted: item.IsCompleted || false,
	}

	c.JSON(http.StatusCreated, todos[stId])
}

func generateId() string {
	rand.NewSource(time.Now().UnixNano())

	randId := rand.Intn(1000)

	return strconv.Itoa(randId)
}

func validateDuplicateInput(item *Todo) error {

	for _, v := range todos {
		if v.Todo == item.Todo {
			return fmt.Errorf("this todo (%s) already exists", item.Todo)
		}
	}
	return nil
}
