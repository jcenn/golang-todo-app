package controllers

import (
	"database/sql"
	"todo-list/internal/todos"
)

var todoService todos.TodoService

func InitializeServices(db *sql.DB){
	todoService = *todos.NewTodoService(db)
}