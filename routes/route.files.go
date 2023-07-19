package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/imsujan276/go-clean-repo/controllers/file-controllers"
	"github.com/imsujan276/go-clean-repo/handlers/file-handlers"
	"github.com/imsujan276/go-clean-repo/middlewares"
)

func InitFileRoutes(db *gorm.DB, route *gin.RouterGroup) {

	fileRepository := filecontrollers.NewFileRepository(db)
	fileService := filecontrollers.NewFileService(fileRepository)
	fileHanlders := fileHandlers.NewCreateHandler(fileService)

	// added auth middlewares
	route.Use(middlewares.Auth())

	route.POST("/create", fileHanlders.CreateHandler)

	route.GET("/", fileHanlders.GetAllFilesHandler)

	route.DELETE("/:fileId", fileHanlders.DeleteHandler)
}
