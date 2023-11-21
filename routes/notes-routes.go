package routes

import (
	"notes-appapi/controllers"
	"notes-appapi/middleware"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(r *gin.Engine) {
	notesRoutes := r.Group("/notes")
	notesRoutes.Use(middleware.AuthorizationWithTokenParsing())

	notesRoutes.GET("/", controllers.GetUsersNotes())
	notesRoutes.POST("/", controllers.AddNote())
	notesRoutes.GET("/recent", controllers.GetUsersRecentNotes())
	notesRoutes.GET("/favorites", controllers.GetUsersFavoriteNotes())
	notesRoutes.PUT("/favorite", controllers.UpdateNotesFavorite())
	notesRoutes.POST("/update", controllers.UpdateNote())
	notesRoutes.GET("/id", controllers.GetNoteById())
}
