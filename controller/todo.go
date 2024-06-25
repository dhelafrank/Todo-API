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
		return
	}
	data.Todos = append(data.Todos, newTodo)
	Context.IndentedJSON(http.StatusOK, newTodo)
}
