package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Todo struct {
	ID          string `json:"id" form:"id"`
	Todo        string `json:"todo" form:"todo"`
	IsCompleted bool   `json:"completed" form:"completed"`
}

// Temporary data source
var Todos = []Todo{
	{ID: "1", Todo: "Read the Book", IsCompleted: false},
	{ID: "2", Todo: "Go to Market", IsCompleted: true},
}

func GenerateId() string {
	rand.NewSource(time.Now().UnixNano())
	randId := rand.Intn(1000)
	return strconv.Itoa(randId)
}

func ValidateDuplicateInput(item *Todo) error {
	for _, v := range Todos {
		if v.Todo == item.Todo {
			return fmt.Errorf("this todo (%s) already exists", item.Todo)
		}
	}
	return nil
}
