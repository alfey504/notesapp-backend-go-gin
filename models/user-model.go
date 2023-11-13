package models

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var hashCost = 14

type User struct {
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UserWithId struct {
	Id       *primitive.ObjectID `json:"_id,omitempty"`
	Username string              `json:"username,omitempty" validate:"required"`
	Password string              `json:"password,omitempty" validate:"required"`
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
