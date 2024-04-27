package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/hot-oranges-86/Todo/models"

	"github.com/hot-oranges-86/Todo/database"
)

func CreateTodo(c echo.Context) error {
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createdTodo, err := database.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, createdTodo)
}
