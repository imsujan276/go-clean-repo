package register

import (
	"github.com/imsujan276/go-clean-repo/controllers/auth-controllers/register"
	"github.com/imsujan276/go-clean-repo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service register.Service
}

func NewHandlerRegister(service register.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input register.RegisterInput
	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	utils.APIResponse(ctx, "Invalid Data", http.StatusBadRequest, nil)
	// 	return
	// }

	if err := ctx.ShouldBind(&input); err != nil {
		utils.APIResponse(ctx, "Invalid Data", http.StatusBadRequest, nil)
		return
	}

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Username",
				Message: "Username is required ",
			},
			{
				Tag:     "lowercase",
				Field:   "Username",
				Message: "Username must be using lowercase",
			},
			{
				Tag:     "required",
				Field:   "Email",
				Message: "Email is required ",
			},
			{
				Tag:     "email",
				Field:   "Email",
				Message: "Email format is not valid",
			},
			{
				Tag:     "required",
				Field:   "Password",
				Message: "Password is required ",
			},
			{
				Tag:     "gte",
				Field:   "Password",
				Message: "Password minimum must be 8 character",
			},
		},
	}

	errorResponse, errCount := utils.GoValidator(&input, config.Options)
	if errCount > 0 {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, errorResponse)
		return
	}

	registerResult, errorCode := h.service.RegisterService(&input)

	switch errorCode {
	case http.StatusCreated:
		accessTokenData := map[string]interface{}{"id": registerResult.ID, "email": registerResult.Email}
		accessToken, errToken := utils.Sign(accessTokenData, utils.GodotEnv("JWT_SECRET"), 60)

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, nil)
			return
		}
		//  parse the UserEntity into response json
		var data register.RegisterResponse

		utils.ObjectToJson(registerResult, &data)

		data.Token = accessToken

		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, data)
		return

	case http.StatusConflict:
		utils.APIResponse(ctx, "Username/Email already taken", http.StatusConflict, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to create an account", http.StatusExpectationFailed, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, nil)
	}

}
