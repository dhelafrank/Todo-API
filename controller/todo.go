package controller

import (
	"net/http"
	"todo-api/data"
	"todo-api/types"

	"github.com/gin-gonic/gin"
)

func GetTodos(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, data.Todos)
}

func AddTodo(Context *gin.Context) {
	var newTodo types.Todo
	if err := Context.BindJSON(&newTodo); err != nil {
		Context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data Type"})
		return
	}
	data.Todos = append(data.Todos, newTodo)
	Context.IndentedJSON(http.StatusOK, newTodo)
}

func GetTodo(Context *gin.Context) {
	id := Context.Param("id")
	todoGotten, error := data.FetchTodoById(id)
	if error != nil {
		Context.JSON(http.StatusNotFound, gin.H{"error": "Todo Not Found"}) //Respond with an error object if error is not nil
		return
	}
	Context.IndentedJSON(http.StatusOK, &todoGotten)
}

func ToggleTodoStatus(Context *gin.Context) {
	id := Context.Param("id")
	todoGotten, error := data.FetchTodoById(id)
	if error != nil {
		Context.JSON(http.StatusNotFound, gin.H{"error": "Todo Not Found"}) //Respond with an error object if error is not nil
		return
	}
	todoGotten.Completed = !todoGotten.Completed
	Context.IndentedJSON(http.StatusOK, &todoGotten)

}
