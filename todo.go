package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Go to the gym", Completed: false},
	{ID: "3", Item: "Watch a movie", Completed: false},
	{ID: "4", Item: "Cook Lunch", Completed: false},
	{ID: "5", Item: "Take a shower", Completed: false},
	{ID: "6", Item: "Code", Completed: false},
	{ID: "7", Item: "Take a nap", Completed: false},
}

// GET REQUEST
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// POST REQUEST
func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// FUNCTION THAT THE HANDLER IS GOING TO UTILISE
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

// GET BY ID REQUEST
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

// PATCH/UPDATE REQUEST
func toggleToDoStatus(context *gin.Context) {
	id := context.Param("id")

	// Find the todo item by ID
	for i, t := range todos {
		if t.ID == id {
			// Toggle the completed status
			todos[i].Completed = !todos[i].Completed
			context.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// DELETE REQUEST
func deleteTodo(context *gin.Context) {
	id := context.Param("id")

	// Find and delete the todo item
	for i, t := range todos {
		if t.ID == id {
			// Remove the item from the slice
			todos = append(todos[:i], todos[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
			return
		}
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleToDoStatus)
	router.DELETE("/todos/:id", deleteTodo) // DELETE endpoint added
	router.Run("localhost:9090")
}
