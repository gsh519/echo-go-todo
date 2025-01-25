package requests

type UpdateTodoRequest struct {
	Content string `json:"content" validate:"required"`
}
