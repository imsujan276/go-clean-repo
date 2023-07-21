package todocontrollers

import (
	"github.com/imsujan276/go-clean-repo/models"
)

type Service interface {
	// create todo and return the todo and status code
	CreateTodo(input *TodoInput) (*models.TodoEntity, int)

	// get all todos and return the todos and status code
	GetAllTodos(userId uint) ([]models.TodoEntity, int)

	// Get todo by id amd return the todo and status code
	GetTodoById(todoId uint) (*models.TodoEntity, int)

	// Update todo by id and return the todo and status code
	UpdateTodoById(input *TodoInput) (*models.TodoEntity, int)

	// Update todo status and return the todo and status code
	UpdateTodoStatus(input *TodoStatusInput) (*models.TodoEntity, int)

	// Delete todo by id and return the status code
	DeleteTodoById(todoId uint) int
}

type service struct {
	repository Repository
}

func NewTodoService(r Repository) *service {
	return &service{repository: r}
}

func (s *service) CreateTodo(input *TodoInput) (*models.TodoEntity, int) {
	todoModel := models.TodoEntity{
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.UserId,
		Completed:   input.Completed,
	}
	return s.repository.CreateTodo(&todoModel)
}

func (s *service) GetAllTodos(userId uint) ([]models.TodoEntity, int) {
	return s.repository.GetAllTodos(userId)
}

func (s *service) GetTodoById(todoId uint) (*models.TodoEntity, int) {
	return s.repository.GetTodoById(todoId)
}

func (s *service) UpdateTodoById(input *TodoInput) (*models.TodoEntity, int) {
	todoModel := models.TodoEntity{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		UserID:      input.UserId,
		Completed:   input.Completed,
	}
	return s.repository.UpdateTodoById(&todoModel)
}

func (s *service) DeleteTodoById(todoId uint) int {
	return s.repository.DeleteTodoById(todoId)
}

func (s *service) UpdateTodoStatus(input *TodoStatusInput) (*models.TodoEntity, int) {
	todoModel := models.TodoEntity{
		ID:        input.ID,
		Completed: input.Completed,
		UserID:    input.UserId,
	}
	return s.repository.UpdateTodoStatus(&todoModel)
}
