package main

import (
	"log"

	"github.com/labstack/echo/v4"

	"github.com/hot-oranges-86/Todo/database"
	"github.com/hot-oranges-86/Todo/handlers"
)

func main() {
	log.Println("123 go")

	e := echo.New()

	err := database.ConnectToPostgreSQL()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	e.POST("/todos", handlers.CreateTodo)

}
