package auth_handler

import (
	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"
	auth_controller "library-service/controllers/auth-controller"
	"library-service/utils"
	"net/http"
)

func (h *authHandler) LoginHandler(ctx *gin.Context) {
	var input auth_controller.UserInput
	ctx.ShouldBindJSON(&input)

	config := gpc.ErrorConfig{
		Options: []gpc.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Username",
				Message: "username is required on body",
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
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}
	resultLogin, errLogin := h.service.LogInUser(&input)

	switch errLogin {

	case http.StatusNotFound:
		utils.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return

	case http.StatusUnauthorized:
		utils.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return

	case http.StatusAccepted:
		jwtToken := utils.GetJwtToken(resultLogin.Username)
		utils.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, jwtToken)
		return

	default:
		utils.APIResponse(ctx, "Unknown error occured", http.StatusInternalServerError, http.MethodPost, nil)
	}
}
