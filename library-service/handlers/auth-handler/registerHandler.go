package auth_handler

import (
	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	auth_controller "library-service/controllers/auth-controller"
	"library-service/utils"
	"net/http"
)

func (h *authHandler) RegisterHandler(ctx *gin.Context) {
	var input auth_controller.UserInput
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Username",
				Message: "Username is required on body",
			},
			{
				Tag:     "lowercase",
				Field:   "Username",
				Message: "Username must be using lowercase",
			},
			{
				Tag:     "required",
				Field:   "Password",
				Message: "Password is required on body",
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
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errorResponse)
		return
	}
	registerResult, errorCode := h.service.RegisterUser(&input)

	switch errorCode {
	case http.StatusCreated:
		jwtToken := utils.GetJwtToken(registerResult.Username)

		utils.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, jwtToken)
		return

	case http.StatusConflict:
		utils.APIResponse(ctx, "Username already taken", http.StatusConflict, http.MethodPost, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to create an account", http.StatusExpectationFailed, http.MethodPost, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, http.MethodPost, nil)
	}
}
