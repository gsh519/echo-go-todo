package main

import (
	"github.com/gsh519/echo-go-todo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/todo", handlers.FetchTodoHandler)
	e.POST("/todo", handlers.CreateTodoHandler)
	e.PUT("/todo/:id", handlers.UpdateTodoHandler)
	e.PUT("/todo/:id/done", handlers.DoneTodoHandler)
	e.DELETE("/todo/:id", handlers.DeleteTotoHandler)
	e.GET("/todo/done", handlers.FetchDoneTodoHandler)

	e.Logger.Fatal(e.Start(":8888"))
}
