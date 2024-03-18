package auth_controller

type UserInput struct {
	Username string `json:"username" validate:"required,lowercase"`
	Password string `json:"password" validate:"required,gte=8"`
}
