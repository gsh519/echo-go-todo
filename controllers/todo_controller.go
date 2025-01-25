package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gsh519/echo-go-todo/models"
	"github.com/gsh519/echo-go-todo/requests"
	"github.com/gsh519/echo-go-todo/services"
	"github.com/gsh519/echo-go-todo/validators"
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
		todos, err := tc.service.FetchTodoService()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "TODOの取得に失敗しました。時間をあけてもう一度お試しください。")
		}

		return c.JSON(http.StatusOK, todos)
	}
}

func (tc *TodoController) CreateTodoHandler(e *echo.Echo) echo.HandlerFunc {
	e.Validator = &validators.Validator{Validator: validator.New()}

	return func(c echo.Context) error {
		payload := new(models.Todo)

		if err := c.Bind(payload); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		if err := c.Validate(payload); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		todo, err := tc.service.CreateTodoService(payload)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc *TodoController) UpdateTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		req := new(requests.UpdateTodoRequest)
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		err := tc.service.UpdateTodoService(id, req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]bool{"success": true})
	}
}

func (tc *TodoController) DoneTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := tc.service.DoneTodoService(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]bool{"success": true})
	}
}

func (tc *TodoController) DeleteTotoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		err := tc.service.DeleteTodoService(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]bool{"success": false})
		}

		return c.JSON(http.StatusOK, map[string]bool{"success": true})
	}
}

func (tc *TodoController) FetchDoneTodoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		dones, err := tc.service.FetchDoneTodoService()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, dones)
	}
}
