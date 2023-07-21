package todocontrollers

import (
	"github.com/imsujan276/go-clean-repo/models"
)

type Service interface {
	CreateTodo(input *TodoInput) (*models.TodoEntity, int)
	GetAllTodos(userId uint) ([]models.TodoEntity, int)
	GetTodoById(todoId uint) (*models.TodoEntity, int)
	UpdateTodoById(input *TodoInput) (*models.TodoEntity, int)
	UpdateTodoStatus(input *TodoStatusInput) (*models.TodoEntity, int)
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
