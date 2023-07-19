package loginHandler

import (
	loginAuth "github.com/imsujan276/go-clean-repo/controllers/auth-controllers/login"
	"github.com/imsujan276/go-clean-repo/controllers/auth-controllers/register"
	"github.com/imsujan276/go-clean-repo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input loginAuth.LoginInput
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			{
				Tag:     "email",
				Field:   "Email",
				Message: "email format is not valid",
			},
			{
				Tag:     "required",
				Field:   "Password",
				Message: "password is required on body",
			},
		},
	}
	errResponse, errCount := utils.GoValidator(&input, config.Options)

	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, errResponse)
		return
	}
	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {

	case http.StatusNotFound:
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, nil)
		return

	case http.StatusUnauthorized:
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, nil)
		return

	case http.StatusAccepted:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := utils.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		var data register.RegisterResponse

		utils.ObjectToJson(resultLogin, &data)

		data.Token = accessToken

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, nil)
			return
		}
		utils.APIResponse(ctx, "Login successfully", http.StatusOK, data)
		return

	default:

		utils.APIResponse(ctx, "Unknown error occured", http.StatusInternalServerError, nil)
	}
}
