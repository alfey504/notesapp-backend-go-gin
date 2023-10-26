package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func VerifyAndParseRequest[T interface{}](c *gin.Context) (*T, error) {

	var model T
	if err := c.BindJSON(&model); err != nil {
		return nil, err
	}

	if err := validate.Struct(&model); err != nil {
		return nil, err
	}

	return &model, nil
}
