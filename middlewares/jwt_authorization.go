package middlewares

import (
	"net/http"

	"github.com/imsujan276/go-clean-repo/utils"

	"github.com/gin-gonic/gin"
)

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Method  string `json:"method"`
	Message string `json:"message"`
}

func Auth() gin.HandlerFunc {

	return gin.HandlerFunc(func(ctx *gin.Context) {

		var errorResponse UnathorizatedError

		if ctx.GetHeader("Authorization") == "" {
			errorResponse.Status = "Forbidden"
			errorResponse.Code = http.StatusForbidden
			errorResponse.Method = ctx.Request.Method
			errorResponse.Message = "Unathorizated Access"
			ctx.JSON(http.StatusForbidden, errorResponse)
			defer ctx.AbortWithStatus(http.StatusForbidden)
		}

		token, err := utils.VerifyTokenHeader(ctx, "JWT_SECRET")

		if err != nil {
			errorResponse.Status = "Unathorizated"
			errorResponse.Code = http.StatusUnauthorized
			errorResponse.Method = ctx.Request.Method
			errorResponse.Message = "Invalid Access token"
			ctx.JSON(http.StatusUnauthorized, errorResponse)
			defer ctx.AbortWithStatus(http.StatusUnauthorized)
		} else {
			// global value result
			ctx.Set("user", token.Claims)
			// return to next method if token is exist
			ctx.Next()
		}
	})
}
