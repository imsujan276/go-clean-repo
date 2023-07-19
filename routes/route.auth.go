package routes

import (
	"github.com/imsujan276/go-clean-repo/controllers/auth-controllers/login"
	"github.com/imsujan276/go-clean-repo/controllers/auth-controllers/register"
	"github.com/imsujan276/go-clean-repo/handlers/auth-handlers/login"
	registerHandler "github.com/imsujan276/go-clean-repo/handlers/auth-handlers/register"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.RouterGroup) {
	loginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(loginRepository)
	loginHandler := loginHandler.NewHandlerLogin(loginService)

	registerRepository := register.NewRegisterRepository(db)
	registerService := register.NewRegisterService(registerRepository)
	registerHandlers := registerHandler.NewHandlerRegister(registerService)

	route.POST("/login", loginHandler.LoginHandler)

	route.POST("/register", registerHandlers.RegisterHandler)
}
