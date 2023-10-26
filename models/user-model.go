package models

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var hashCost = 14

type User struct {
	UserId   Oid    `json:"_id,omitempty"`
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type Oid struct {
	Id string `json:"$oid,omitempty"`
}

var validate = validator.New()

func ValidateUser(c *gin.Context) (*User, error) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return nil, err
	}

	if err := validate.Struct(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (user *User) HashPassword() error {
	tempPassword := user.Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(tempPassword), hashCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user User) CompareHash(userTwo User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userTwo.Password))
}
