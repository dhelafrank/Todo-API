package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default() //Here router is our server
	
	//ROUTES
	router.GET("/todos") 
	
	router.Run("localhost:9090")
}
