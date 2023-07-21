package todohandlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imsujan276/go-clean-repo/utils"
)

func (h *handler) DeleteTodoById(context *gin.Context) {
	todoIdParam := context.Param("todoId")
	todoIdInt, err := (strconv.ParseInt(todoIdParam, 0, 0))
	todoId := uint(todoIdInt)
	if err != nil {
		utils.APIResponse(context, "Cannot parse the parameter", http.StatusBadRequest, nil)
	}

	statusCode := h.service.DeleteTodoById(todoId)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(context, "Success", http.StatusOK, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return
	case http.StatusNotFound:
		utils.APIResponse(context, "No Todo Found", http.StatusNotFound, nil)
		return
	}

}
