package todohandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todocontrollers "github.com/imsujan276/go-clean-repo/controllers/todo-controllers"
	"github.com/imsujan276/go-clean-repo/models"
	"github.com/imsujan276/go-clean-repo/utils"
)

func (h *handler) CreateTodo(context *gin.Context) {
	var user models.UserEntity
	var todoInput todocontrollers.TodoInput

	jwtData, _ := context.Get("user")
	errors := utils.StringToEntity(jwtData, &user)
	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, nil)
		return
	}

	if err := context.ShouldBindJSON(&todoInput); err != nil {
		utils.APIResponse(context, "Invalid Data", http.StatusBadRequest, nil)
		return
	}

	todoInput.UserId = user.ID

	todoResponse, statusCode := h.service.CreateTodo(&todoInput)
	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Todo created successfully.", http.StatusCreated, todoResponse)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return
	}

}
