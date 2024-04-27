package handlers

import (
	"net/http"
	"strconv"

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

func GetTodos(c echo.Context) error {
	todos, err := database.GetTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	todo, err := database.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var updatedTodo models.Todo
	if err := c.Bind(&updatedTodo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	updatedTodo.ID = id

	result, err := database.UpdateTodo(updatedTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteTodo(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	err = database.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return (c.NoContent(http.StatusNoContent))
}
