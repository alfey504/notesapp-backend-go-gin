package middleware

import (
	"net/http"
	"notes-appapi/responses"
	"notes-appapi/utils"

	"github.com/gin-gonic/gin"
)

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
