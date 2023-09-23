package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"todo-list/internal/todos"
)

func AddTodoHandler(c echo.Context) error {
	var request todos.AddTodoRequest
	if err := c.Bind(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	todo := todoService.AddTodo(request)

	return c.JSON(http.StatusCreated, todo)
}

func GetTodosHandler(c echo.Context) error {
	return c.JSON(http.StatusAccepted, todoService.GetTodos())
}

func GetTodoByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusAccepted, todoService.GetTodoById(uint(id)))
}

func EditTodoHandler(c echo.Context) error {
	id, _:= strconv.Atoi(c.Param("id"))
	
	var req todos.EditTodoRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusAccepted, todoService.EditTodo(uint(id), req))
}

func DeleteTodoHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "id parameter has to be an uint")
	}
	todoService.DeleteTodo(uint(id))
	return c.String(http.StatusAccepted, "")
}

func DeleteAllTodosHandler(c echo.Context) error {
	todoService.DeleteAllTodos()
	return c.String(http.StatusAccepted, "")
}