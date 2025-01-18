package handlers

import (
	"net/http"
	"time"

	"github.com/gsh519/echo-go-todo/models"
	"github.com/labstack/echo/v4"
)

func FetchTodoHandler(c echo.Context) error {
	todos := []models.Todo{
		{
			TodoID:      1,
			Content:     "hoge",
			IsDone:      false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedFlag: false,
		},
		{
			TodoID:      2,
			Content:     "fuga",
			IsDone:      false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedFlag: false,
		},
		{
			TodoID:      3,
			Content:     "piyo",
			IsDone:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedFlag: false,
		},
		{
			TodoID:      4,
			Content:     "good",
			IsDone:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedFlag: true,
		},
	}

	return c.JSON(http.StatusOK, todos)
}

func CreateTodoHandler(c echo.Context) error {
	todo := models.Todo{
		TodoID:      5,
		Content:     "bad",
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedFlag: false,
	}

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodoHandler(c echo.Context) error {
	todo := models.Todo{
		TodoID:      5,
		Content:     "hogefuga",
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedFlag: false,
	}

	return c.JSON(http.StatusOK, todo)
}

func DoneTodoHandler(c echo.Context) error {
	todo := models.Todo{
		TodoID:      5,
		Content:     "hogefuga",
		IsDone:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedFlag: false,
	}

	return c.JSON(http.StatusOK, todo)
}

func DeleteTotoHandler(c echo.Context) error {
	type Response struct {
		Success bool `json:"success"`
	}

	var isSuccess bool

	return c.JSON(http.StatusOK, Response{Success: isSuccess})
}

func FetchDoneTodoHandler(c echo.Context) error {
	dones := []models.Todo{
		{
			TodoID:      3,
			Content:     "piyo",
			IsDone:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedFlag: false,
		},
		{
			TodoID:      4,
			Content:     "good",
			IsDone:      true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedFlag: true,
		},
	}

	return c.JSON(http.StatusOK, dones)
}
