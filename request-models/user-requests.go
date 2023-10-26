package request_models

import "notes-appapi/models"

type LoginAndSignUpRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func (loginAndSignUpRequest LoginAndSignUpRequest) GetUserModel() models.User {
	return models.User{
		Username: loginAndSignUpRequest.Username,
		Password: loginAndSignUpRequest.Password,
	}
}
