package register

import (
	"github.com/imsujan276/go-clean-repo/models"
	"github.com/imsujan276/go-clean-repo/utils"
	"net/http"
)

type Service interface {
	RegisterService(*RegisterInput) (*models.UserEntity, int)
}

type service struct {
	repository Repository
}

func NewRegisterService(repository *repository) *service {
	return &service{repository: repository}
}

func (service *service) RegisterService(input *RegisterInput) (*models.UserEntity, int) {
	fileName := ""
	if input.Image != nil {
		fname, err := utils.UploadedFormDataImg(input.Image)
		if err != nil {
			return nil, http.StatusInternalServerError
		}
		fileName = fname
	}

	user := models.UserEntity{
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
		Image:    fileName,
	}
	return service.repository.RegisterRepository(&user)
}
