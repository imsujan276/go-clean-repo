package todocontrollers

import (
	"net/http"

	"github.com/imsujan276/go-clean-repo/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	// create todo and return the todo and status code
	CreateTodo(todo *models.TodoEntity) (*models.TodoEntity, int)

	// get all todos and return the todos and status code
	GetAllTodos(userId uint) ([]models.TodoEntity, int)

	// Get todo by id amd return the todo and status code
	GetTodoById(todoId uint) (*models.TodoEntity, int)

	// Update todo by id and return the todo and status code
	UpdateTodoById(todo *models.TodoEntity) (*models.TodoEntity, int)

	// Update todo status and return the todo and status code
	UpdateTodoStatus(todo *models.TodoEntity) (*models.TodoEntity, int)

	// Delete todo by id and return the status code
	DeleteTodoById(todoId uint) int
}

type repository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) CreateTodo(todo *models.TodoEntity) (*models.TodoEntity, int) {
	db := repo.db

	db.NewRecord(todo)
	createTodo := db.Create(&todo)

	if createTodo.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return todo, http.StatusCreated
}

func (repo *repository) GetAllTodos(userId uint) ([]models.TodoEntity, int) {
	db := repo.db
	var todos []models.TodoEntity

	checkIfFileExists := db.Select("*").Where("user_id=?", userId).Find(&todos)
	// db.Preload will populate the user field in `TodosEntity`
	db.Preload("User").Find(&todos)
	if checkIfFileExists.Error != nil {
		return nil, http.StatusNotFound
	}
	return todos, http.StatusOK
}

func (repo *repository) GetTodoById(todoId uint) (*models.TodoEntity, int) {
	db := repo.db
	var todo models.TodoEntity

	checkIfFileExists := db.Where("id = ?", todoId).Find(&todo)
	db.Preload("User").Find(&todo)
	if checkIfFileExists.Error != nil {
		return nil, http.StatusNotFound
	}

	return &todo, http.StatusOK
}

func (repo *repository) UpdateTodoById(todo *models.TodoEntity) (*models.TodoEntity, int) {
	db := repo.db
	existingTodo := &models.TodoEntity{}

	checkIfFileExists := db.Where("id = ?", todo.ID).Where("user_id = ?", todo.UserID).Find(&existingTodo)
	if checkIfFileExists.Error != nil {
		return nil, http.StatusNotFound
	}

	// Update only the relevant fields of the existing todo with the new values
	existingTodo.Title = todo.Title
	existingTodo.Description = todo.Description
	existingTodo.Completed = todo.Completed
	db.Save(existingTodo)

	return existingTodo, http.StatusOK
}

func (repo *repository) UpdateTodoStatus(todo *models.TodoEntity) (*models.TodoEntity, int) {
	db := repo.db
	existingTodo := &models.TodoEntity{}

	checkIfFileExists := db.Where("id = ?", todo.ID).Where("user_id = ?", todo.UserID).Find(&existingTodo)
	if checkIfFileExists.Error != nil {
		return nil, http.StatusNotFound
	}
	existingTodo.Completed = todo.Completed
	db.Save(existingTodo)
	return existingTodo, http.StatusOK
}

func (repo *repository) DeleteTodoById(todoId uint) int {
	db := repo.db
	var todo models.TodoEntity

	checkIfFileExists := db.Select("*").Where("id=?", todoId).Find(&todo)

	if checkIfFileExists.RowsAffected > 0 {
		db.Delete(&todo)
		return http.StatusOK
	}
	return http.StatusNotFound
}
