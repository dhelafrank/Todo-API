package controller

import (
	"net/http"
	"todo-api/data"

	"github.com/gin-gonic/gin"
)

func GetTodos(Context *gin.Context) {
	Context.IndentedJSON(http.StatusOK, data.Todos)
}
