package todohandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todocontrollers "github.com/imsujan276/go-clean-repo/controllers/todo-controllers"
	"github.com/imsujan276/go-clean-repo/models"
	"github.com/imsujan276/go-clean-repo/utils"
)

func (h *handler) UpdateTodoById(context *gin.Context) {
	todoIdParam := context.Param("todoId")
	todoIdInt, err := (strconv.ParseInt(todoIdParam, 0, 0))
	todoId := uint(todoIdInt)
	if err != nil {
		utils.APIResponse(context, "Cannot parse the parameter", http.StatusBadRequest, nil)
	}

	var user models.UserEntity
	jwtData, _ := context.Get("user")
	errors := utils.StringToEntity(jwtData, &user)
	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, nil)
		return
	}

	var todoInput todocontrollers.TodoInput
	if err := context.ShouldBindJSON(&todoInput); err != nil {
		utils.APIResponse(context, "Invalid Data", http.StatusBadRequest, nil)
		return
	}

	todoInput.UserId = user.ID
	todoInput.ID = todoId

	todosResponse, statusCode := h.service.UpdateTodoById(&todoInput)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Success", http.StatusOK, &todosResponse)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return
	case http.StatusNotFound:
		utils.APIResponse(context, "No Todo Found", http.StatusNotFound, nil)
		return
	}

}

func (h *handler) UpdateTodoStatus(context *gin.Context) {
	var todoStatusInput todocontrollers.TodoStatusInput
	if err := context.ShouldBindJSON(&todoStatusInput); err != nil {
		utils.APIResponse(context, "Invalid Data", http.StatusBadRequest, nil)
		return
	}

	var user models.UserEntity
	jwtData, _ := context.Get("user")
	errors := utils.StringToEntity(jwtData, &user)
	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, nil)
		return
	}
	todoStatusInput.UserId = user.ID

	todosResponse, statusCode := h.service.UpdateTodoStatus(&todoStatusInput)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Success", http.StatusOK, &todosResponse)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return
	case http.StatusNotFound:
		utils.APIResponse(context, "No Todo Found", http.StatusNotFound, nil)
		return
	}

}
