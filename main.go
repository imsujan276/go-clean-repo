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
	file := api.Group("/file")

	routes.InitAuthRoutes(db, api)
	routes.InitFileRoutes(db, file)

	SetupStaticFiles(router)
	return router
}

func SetupStaticFiles(router *gin.Engine) {
	// Serve static files from the "uploads/images" directory
	router.Static("/uploads/images", "./uploads/images")
	router.Static("/uploads/files", "./uploads/files")
}
