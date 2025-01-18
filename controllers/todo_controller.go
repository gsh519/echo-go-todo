package controllers

import (
	"net/http"

	"github.com/gsh519/echo-go-todo/services"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	service *services.AppService
}

func NewTodoController(service *services.AppService) *TodoController {
	return &TodoController{service: service}
}

func (tc *TodoController) FetchTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		todos := tc.service.FetchTodoService()

		return c.JSON(http.StatusOK, todos)
	}
}

func (tc *TodoController) CreateTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := tc.service.CreateTodoService()

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc *TodoController) UpdateTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := tc.service.UpdateTodoService()

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc *TodoController) DoneTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := tc.service.DoneTodoService()

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc *TodoController) DeleteTotoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		type Response struct {
			Success bool `json:"success"`
		}

		isSuccess := tc.service.DeleteTodoService()

		return c.JSON(http.StatusOK, Response{Success: isSuccess})
	}
}

func (tc *TodoController) FetchDoneTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		dones := tc.service.FetchDoneTodoService()

		return c.JSON(http.StatusOK, dones)
	}
}
