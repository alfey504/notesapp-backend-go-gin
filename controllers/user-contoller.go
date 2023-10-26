package controllers

import (
	"context"
	"net/http"
	"notes-appapi/config"
	"notes-appapi/models"
	request_models "notes-appapi/request-models"
	"notes-appapi/responses"
	"notes-appapi/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var userCollection = config.GetCollection(config.DB, "users")

func AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the json some of the fields might be missing",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		if err := validate.Struct(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing json some fields might be missing or incorrect",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		if err := user.HashPassword(); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error while adding user to the database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
		}

		result, err := userCollection.InsertOne(ctx, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error while adding user to the database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Added user successfully",
			Data: map[string]interface{}{
				"data": result,
			},
		})
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// validate and parse user from request
		userPointer, err := utils.VerifyAndParseRequest[request_models.LoginAndSignUpRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing json some fields might be missing or incorrect",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		userInRequest := *userPointer
		user := userInRequest.GetUserModel()

		// getting user from the database and checking if user exists
		result := userCollection.FindOne(ctx, bson.M{"username": user.Username})
		if result.Err() == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "Incorrect username or password",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		} else if result.Err() != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error getting the user from database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		// decoding user fetched from the database
		var existingUser models.User
		if err := result.Decode(&existingUser); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "Error logging in the user",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		// verifying the password
		if err := existingUser.CompareHash(user); err != nil {
			c.JSON(http.StatusOK, responses.Response{
				Status:  http.StatusConflict,
				Message: "Incorrect username or password",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		token, err := utils.GenerateJwtToken(existingUser.Username)

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Successfully logged in user",
			Data: map[string]interface{}{
				"data": map[string]interface{}{
					"username": existingUser.Username,
					"token":    token,
				},
			},
		})
	}
}

func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Auth Successful",
			Data: map[string]interface{}{
				"data": "success",
			},
		})
	}
}
