package main

import (
	"log"

	"github.com/imsujan276/go-clean-repo/configs"
	"github.com/imsujan276/go-clean-repo/routes"
	"github.com/imsujan276/go-clean-repo/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupAppRouter()
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func SetupAppRouter() *gin.Engine {

	service := configs.NewDBService()
	db := service.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	api := router.Group("api/v1")
	routes.InitAuthRoutes(db, api)

	file := api.Group("/file")
	routes.InitFileRoutes(db, file)

	todo := api.Group("/todo")
	routes.InitTodoRoutes(db, todo)

	SetupStaticFiles(router)
	return router
}

func SetupStaticFiles(router *gin.Engine) {
	router.Static("/uploads/images", "./uploads/images")
	router.Static("/uploads/files", "./uploads/files")
}
