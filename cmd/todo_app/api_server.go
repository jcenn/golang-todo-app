package main

import (
	"todo-list/internal/controllers"
	"todo-list/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database := db.ConnectToDB()
	controllers.InitializeServices(database)
	echo := echo.New()
	echo.Use(middleware.CORS())
	echo.GET("/", controllers.IndexHandler)
	echo.POST("/todos/add", controllers.AddTodoHandler)
	echo.GET("/todos", controllers.GetTodosHandler)
	echo.GET("/todos/:id", controllers.GetTodoByIdHandler)
	echo.PATCH("/todos/:id", controllers.EditTodoHandler)
	echo.DELETE("/todos/:id", controllers.DeleteTodoHandler)
	echo.DELETE("/todos", controllers.DeleteAllTodosHandler)
	echo.Logger.Fatal(echo.Start(":8080"))
}
