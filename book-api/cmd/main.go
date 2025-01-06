package main

import (
	"book-api/internal/handlers"
	"book-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitMongoDB()
	r := gin.Default()
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBook)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.Run(":8080")
}
