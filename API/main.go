package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/hot-oranges-86/Todo/database"
	"github.com/hot-oranges-86/Todo/handlers"
	"github.com/hot-oranges-86/Todo/middlewares"
)

func main() {
	log.Println("123 go")

	e := echo.New()

	err := database.ConnectToPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	e.Use(middlewares.GetMiddlewares()...)

	e.POST("/todos", handlers.CreateTodo)
	e.GET("/todos", handlers.GetTodos)
	e.PUT("/todos/:id", handlers.UpdateTodo)
	e.DELETE("/todos/:id", handlers.DeleteTodo)

	log.Println("Server started on port 8080")
	log.Fatal(e.Start(":8080"))
}
