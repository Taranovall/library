package auth_handler

import auth_controller "library-service/controllers/auth-controller"

type authHandler struct {
	service auth_controller.Service
}

func NewAuthHandler(service auth_controller.Service) *authHandler {
	return &authHandler{service: service}
}
