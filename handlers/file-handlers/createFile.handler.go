package fileHandlers

import (
	filecontrollers "github.com/imsujan276/go-clean-repo/controllers/file-controllers"
	"github.com/imsujan276/go-clean-repo/models"
	"github.com/imsujan276/go-clean-repo/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateHandler(context *gin.Context) {
	file, header, _ := context.Request.FormFile("file")

	var user models.UserEntity

	jwtData, _ := context.Get("user")

	// convert header to user enitity
	errors := utils.StringToEntity(jwtData, &user)

	if errors != nil {
		utils.APIResponse(context, "User does not exist", http.StatusNotFound, nil)
		return
	}
	// result, err := utils.UploadFile(file, header.Header.Get("Content-Type"))
	// if err != nil {
	// 	utils.APIResponse(context, "Unable to upload file to the server", http.StatusFailedDependency, nil)
	// }
	// fileInput := filecontrollers.FileInput{
	// 	ID:     result.PublicID,
	// 	Type:   result.Format,
	// 	Name:   header.Filename,
	// 	UserId: user.ID,
	// }

	path, err := utils.UploadImage(file, header, "uploads/files")
	if err != nil {
		utils.APIResponse(context, "Unable to upload file to the server", http.StatusFailedDependency, nil)
	}
	fileInput := filecontrollers.FileInput{
		ID:     utils.GenerateRandomString(10),
		Type:   filepath.Ext(header.Filename),
		Name:   filepath.Base(path),
		Url:    path,
		UserId: user.ID,
	}

	fileResponse, statusCode := h.service.CreateFile(&fileInput)

	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(context, "Uploaded the file successfully.", http.StatusCreated, fileResponse)
		return

	case http.StatusExpectationFailed:
		utils.APIResponse(context, "Internal Server error occured", http.StatusExpectationFailed, nil)
		return

	case http.StatusConflict:
		utils.APIResponse(context, "File already exists. Please try with another file", http.StatusCreated, nil)
		return
	}

}
