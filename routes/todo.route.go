package routes

import (
	"github.com/gin-gonic/gin"
	todocontrollers "github.com/imsujan276/go-clean-repo/controllers/todo-controllers"
	todohandlers "github.com/imsujan276/go-clean-repo/handlers/todo-handlers"
	"github.com/imsujan276/go-clean-repo/middlewares"
	"github.com/jinzhu/gorm"
)

func InitTodoRoutes(db *gorm.DB, route *gin.RouterGroup) {

	todoRepository := todocontrollers.NewTodoRepository(db)
	todoService := todocontrollers.NewTodoService(todoRepository)
	todoHanlders := todohandlers.NewCreateHandler(todoService)

	// added auth middlewares
	route.Use(middlewares.Auth())

	route.POST("/", todoHanlders.CreateTodo)
	route.GET("/", todoHanlders.GetAllTodos)
	route.GET("/:todoId", todoHanlders.GetTodoById)
	route.PUT("/:todoId", todoHanlders.UpdateTodoById)
	route.DELETE("/:todoId", todoHanlders.DeleteTodoById)
	route.POST("/update-status", todoHanlders.UpdateTodoStatus)

	// route.DELETE("/:fileId", fileHanlders.DeleteHandler)
}
