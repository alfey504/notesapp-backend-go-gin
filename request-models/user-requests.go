package request_models

import (
	"notes-appapi/models"

	"github.com/gin-gonic/gin"
)

type LoginAndSignUpRequest struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func (loginAndSignInRequest *LoginAndSignUpRequest) ParseAndValidateRequest(c *gin.Context) error {
	if err := c.BindJSON(&loginAndSignInRequest); err != nil {
		return err
	}
	if err := validate.Struct(&loginAndSignInRequest); err != nil {
		return err
	}
	return nil
}

func (loginAndSignUpRequest LoginAndSignUpRequest) GetUserModel() models.User {
	return models.User{
		Username: loginAndSignUpRequest.Username,
		Password: loginAndSignUpRequest.Password,
	}
}
