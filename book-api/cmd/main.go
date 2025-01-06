package main

import (
	"book-api/internal/handlers"
	"book-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao MongoDB
	repository.InitMongoDB()
	r := gin.Default()
	r.GET("/books", handlers.GetBooks)
	r.Run(":8080")
}
