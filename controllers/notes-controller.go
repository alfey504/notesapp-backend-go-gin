package controllers

import (
	"context"
	"net/http"
	"notes-appapi/config"
	request_models "notes-appapi/request-models"
	"notes-appapi/responses"
	"notes-appapi/services"
	"notes-appapi/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var notesCollection = config.GetCollection(config.DB, "notes")

func AddNote() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var notesRequest request_models.NotesRequest
		if err := notesRequest.ValidateAndParseFromRequest(c); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the json some of the fields might be missing",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		username := c.GetString(config.USERNAME_KEY)

		if username == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the username from request",
				Data: map[string]interface{}{
					"data": "failed to parse username",
				},
			})
			return
		}

		note, err := notesRequest.GetNoteModel(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error parsing the provided date make sure the date is in ISO format",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		result, err := notesCollection.InsertOne(ctx, note)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "failed to save the note",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "successfully saved the note",
			Data: map[string]interface{}{
				"data": result,
			},
		})
	}
}

func GetUsersNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
		defer cancel()

		username := c.GetString(config.USERNAME_KEY)
		if username == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the username from request",
				Data: map[string]interface{}{
					"data": "failed to parse username",
				},
			})
			return
		}

		notes, err := services.GetUsersNotes(ctx, username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error getting data from the database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Successfully fetched data from the database",
			Data: map[string]interface{}{
				"data": notes,
			},
		})
	}
}

func GetUsersRecentNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
		defer cancel()

		username := c.GetString(config.USERNAME_KEY)
		if username == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the username from request",
				Data: map[string]interface{}{
					"data": "failed to parse username",
				},
			})
			return
		}

		notes, err := services.GetUsersRecentNotes(ctx, username)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error getting data from the database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Successfully fetched data from the database",
			Data: map[string]interface{}{
				"data": notes,
			},
		})
	}
}

func GetUsersFavoriteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		username := c.GetString(config.USERNAME_KEY)
		if username == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the username from request",
				Data: map[string]interface{}{
					"data": "failed to parse username",
				},
			})
			return
		}

		notes, err := services.GetUsersFavoriteNotes(ctx, username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error fetching notes from the database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Successfully fetched data from the database",
			Data: map[string]interface{}{
				"data": notes,
			},
		})
	}
}

func UpdateNotesFavorite() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		username := c.GetString(config.USERNAME_KEY)
		if username == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the username from request",
				Data: map[string]interface{}{
					"data": "failed to parse username",
				},
			})
			return
		}

		var notesFavoriteUpdateRequest request_models.NotesFavoriteUpdateRequest
		if err := notesFavoriteUpdateRequest.ValidateAndParseFromRequest(c); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the json some of the fields might be missing",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		favorite, err := utils.ParseBoolFromString(notesFavoriteUpdateRequest.Favorite)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing favorite field",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		err = services.SetNoteFavorite(ctx, notesFavoriteUpdateRequest.Id, favorite)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "error updating notes on the database",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Successfully update favorite",
			Data: map[string]interface{}{
				"data": "successfully updated",
			},
		})
	}
}

func UpdateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		username := c.GetString(config.USERNAME_KEY)
		println("username ", username)
		if username == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing the username from request",
				Data: map[string]interface{}{
					"data": "failed to parse username",
				},
			})
			return
		}

		var noteRequestWithId request_models.NoteRequestWithId
		if err := noteRequestWithId.ValidateAndParseFromRequest(c); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "error parsing request",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		note, err := noteRequestWithId.GetModel(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "Error parsing data from request",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		if err := services.UpdateNote(ctx, note); err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{
				Status:  http.StatusInternalServerError,
				Message: "Error updating data",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Updated Notes successfully",
			Data: map[string]interface{}{
				"data": "success",
			},
		})
	}
}

func GetNoteById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()

		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "missing query Id",
				Data: map[string]interface{}{
					"data": "Please provide id as a query",
				},
			})
			return
		}

		note, err := services.GetNoteById(ctx, id)
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "Unable to find document with given Id",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		} else if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{
				Status:  http.StatusBadRequest,
				Message: "Unable to find document with given Id",
				Data: map[string]interface{}{
					"data": err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusOK, responses.Response{
			Status:  http.StatusOK,
			Message: "Successfully fetched note",
			Data: map[string]interface{}{
				"data": note,
			},
		})
		return
	}
}
