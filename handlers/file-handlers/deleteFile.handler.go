package fileHandlers

import (
	"github.com/imsujan276/go-clean-repo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteHandler(context *gin.Context) {
	fileIdParam := context.Param("fileId")

	fileIdInt, err := (strconv.ParseInt(fileIdParam, 0, 0))

	fileId := uint(fileIdInt)
	if err != nil {
		utils.APIResponse(context, "Cannot parse the parameter", http.StatusBadRequest, nil)
	}

	statusCode := h.service.DeleteFile(fileId)

	switch statusCode {
	case http.StatusOK:

		utils.APIResponse(context, "Deleted file", http.StatusOK, nil)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return

	case http.StatusNotFound:
		utils.APIResponse(context, "File does not exists. Please try with another file", http.StatusNotFound, nil)
		return
	}

}
