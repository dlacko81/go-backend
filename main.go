package main

import (
	"github.com/gin-gonic/gin"
	"go-backend/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/submit", handlers.SubmitData)
	r.GET("/data", handlers.GetData)
	r.Run() // Default on port 8080
}
