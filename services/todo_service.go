package services

import (
	"database/sql"
	"time"

	"github.com/gsh519/echo-go-todo/models"
	"github.com/gsh519/echo-go-todo/requests"
)

type AppService struct {
	db *sql.DB
}

func NewAppService(db *sql.DB) *AppService {
	return &AppService{db: db}
}

func (s *AppService) FetchTodoService() ([]models.Todo, error) {
	const sql = "select * from todos where deleted_flag = false and is_done = false"
	rows, err := s.db.Query(sql)
	if err != nil {
		return []models.Todo{}, err
	}

	todos := make([]models.Todo, 0)

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.TodoID,
			&todo.Content,
			&todo.IsDone,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DeletedFlag,
		)

		if err != nil {
			return []models.Todo{}, err
		} else {
			todos = append(todos, todo)
		}
	}

	return todos, nil
}

func (s *AppService) CreateTodoService(payload *models.Todo) (models.Todo, error) {
	const sql = `
		insert into todos
		(content, is_done, created_at, updated_at, deleted_flag)
		values (?, false, time.Now(), time.Now(), false);
	`

	var newTodo models.Todo
	newTodo.Content, newTodo.IsDone, newTodo.CreatedAt, newTodo.UpdatedAt, newTodo.DeletedFlag = payload.Content, false, time.Now(), time.Now(), false

	result, err := s.db.Exec(sql, payload.Content)
	if err != nil {
		return models.Todo{}, err
	}

	id, _ := result.LastInsertId()
	newTodo.TodoID = int(id)

	return newTodo, nil
}

func (s *AppService) UpdateTodoService(id int, req *requests.UpdateTodoRequest) error {
	const sql = `udpate todos set content = ? where todo_id = ?`

	_, err := s.db.Exec(sql, "content", id, req.Content)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppService) DoneTodoService(id int) error {
	const sql = `update todos set is_done = true;`

	_, err := s.db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppService) DeleteTodoService(id int) error {
	const sql = `delete from todos where todo_id = ?`

	_, err := s.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppService) FetchDoneTodoService() ([]models.Todo, error) {
	const sql = "select * from todos where deleted_flag = false and is_done = true"
	rows, err := s.db.Query(sql)
	if err != nil {
		return []models.Todo{}, err
	}

	todos := make([]models.Todo, 0)

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.TodoID,
			&todo.Content,
			&todo.IsDone,
			&todo.CreatedAt,
			&todo.UpdatedAt,
			&todo.DeletedFlag,
		)

		if err != nil {
			return []models.Todo{}, err
		} else {
			todos = append(todos, todo)
		}
	}

	return todos, nil
}
