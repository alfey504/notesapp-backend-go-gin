package routes

import (
	"notes-appapi/controllers"
	"notes-appapi/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	userRoutes.POST("/signup", controllers.AddUser())
	userRoutes.POST("/login", controllers.LoginUser())

	authUserRoutes := r.Group("/user/auth")
	authUserRoutes.Use(middleware.AuthorizationWithTokenParsing())
	authUserRoutes.GET("/", controllers.VerifyUser())
}
