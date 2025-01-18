package services

import (
	"database/sql"
	"time"

	"github.com/gsh519/echo-go-todo/models"
)

type AppService struct {
	db *sql.DB
}

func NewAppService(db *sql.DB) *AppService {
	return &AppService{db: db}
}

func (s *AppService) FetchTodoService() []models.Todo {
	now := time.Now()
	todos := []models.Todo{
		{
			TodoID:      1,
			Content:     "hoge",
			IsDone:      false,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedFlag: false,
		},
		{
			TodoID:      2,
			Content:     "fuga",
			IsDone:      false,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedFlag: false,
		},
		{
			TodoID:      3,
			Content:     "piyo",
			IsDone:      true,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedFlag: false,
		},
		{
			TodoID:      4,
			Content:     "good",
			IsDone:      true,
			CreatedAt:   now,
			UpdatedAt:   now,
			DeletedFlag: true,
		},
	}

	return todos
}

func (s *AppService) CreateTodoService() models.Todo {
	return models.Todo{
		TodoID:      5,
		Content:     "bad",
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedFlag: false,
	}
}

func (s *AppService) UpdateTodoService() models.Todo {
	return models.Todo{
		TodoID:      5,
		Content:     "bad",
		IsDone:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedFlag: false,
	}
}

func (s *AppService) DoneTodoService() models.Todo {
	return models.Todo{
		TodoID:      5,
		Content:     "hogefuga",
		IsDone:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedFlag: false,
	}
}

func (s *AppService) DeleteTodoService() bool {
	return true
}

func (s *AppService) FetchDoneTodoService() []models.Todo {
	return []models.Todo{
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
}
