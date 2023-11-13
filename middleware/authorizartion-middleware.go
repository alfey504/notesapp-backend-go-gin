package middleware

import (
	"net/http"
	"notes-appapi/config"
	"notes-appapi/responses"
	"notes-appapi/utils"

	"github.com/gin-gonic/gin"
)

const USERNAME_KEY = "username"

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := utils.ValidateToken(c); err != nil {
			c.JSON(http.StatusUnauthorized, responses.Response{
				Status:  http.StatusUnauthorized,
				Message: "User is not authorized",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthorizationWithTokenParsing() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := utils.ValidateTokenAndGetUsername(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, responses.Response{
				Status:  http.StatusUnauthorized,
				Message: "Trouble authenticating user",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			c.Abort()
			return
		}
		c.Set(config.USERNAME_KEY, username)
		c.Next()
	}
}
