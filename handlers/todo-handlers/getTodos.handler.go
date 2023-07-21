package todohandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imsujan276/go-clean-repo/models"
	"github.com/imsujan276/go-clean-repo/utils"
)

func (h *handler) GetAllTodos(context *gin.Context) {
	jwtData, _ := context.Get("user")
	var user models.UserEntity
	errors := utils.StringToEntity(jwtData, &user)

	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, nil)
		return
	}
	userId := user.ID

	todosResponse, statusCode := h.service.GetAllTodos(userId)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Success", http.StatusOK, &todosResponse)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return
	case http.StatusNotFound:
		utils.APIResponse(context, "Todos Not Found", http.StatusNotFound, []models.TodoEntity{})
		return
	}

}

func (h *handler) GetTodoById(context *gin.Context) {
	todoIdParam := context.Param("todoId")
	todoIdInt, err := (strconv.ParseInt(todoIdParam, 0, 0))
	todoId := uint(todoIdInt)
	if err != nil {
		utils.APIResponse(context, "Cannot parse the parameter", http.StatusBadRequest, nil)
	}

	todosResponse, statusCode := h.service.GetTodoById(todoId)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Success", http.StatusOK, &todosResponse)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return
	case http.StatusNotFound:
		utils.APIResponse(context, "Todo Not Found", http.StatusNotFound, nil)
		return
	}

}
