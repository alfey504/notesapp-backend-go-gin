package request_models

import (
	"notes-appapi/models"
	"notes-appapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

// Add notes request
type NotesRequest struct {
	Title    string `json:"title,omitempty" validate:"required"`
	Text     string `json:"text,omitempty" validate:"required"`
	Date     string `json:"date,omitempty" validate:"required"`
	Favorite string `json:"favorite,omitempty" validate:"required"`
}

func (notesRequest *NotesRequest) ValidateAndParseFromRequest(c *gin.Context) error {
	if err := c.BindJSON(notesRequest); err != nil {
		return err
	}
	if err := validate.Struct(notesRequest); err != nil {
		return err
	}

	if _, err := utils.ParseBoolFromString(notesRequest.Favorite); err != nil {
		return err
	}

	return nil
}

func (notesRequest NotesRequest) GetNoteModel(username string) (models.Note, error) {
	date, err := utils.ParseDateFromString(notesRequest.Date)
	if err != nil {
		return models.Note{}, err
	}

	favoriteBool, err := utils.ParseBoolFromString(notesRequest.Favorite)
	if err != nil {
		return models.Note{}, err
	}

	return models.Note{
		Username: username,
		Title:    notesRequest.Title,
		Text:     notesRequest.Text,
		Date:     date,
		Favorite: favoriteBool,
	}, nil
}

// notes favorite update request
type NotesFavoriteUpdateRequest struct {
	Id       string `json:"id,omitempty" validate:"required"`
	Favorite string `json:"favorite,omitempty" validate:"required"`
}

func (notesFavoriteUpdateRequest *NotesFavoriteUpdateRequest) ValidateAndParseFromRequest(c *gin.Context) error {
	if err := c.BindJSON(notesFavoriteUpdateRequest); err != nil {
		return err
	}
	if err := validate.Struct(notesFavoriteUpdateRequest); err != nil {
		return err
	}

	if _, err := utils.ParseBoolFromString(notesFavoriteUpdateRequest.Favorite); err != nil {
		return err
	}
	return nil
}

type NoteRequestWithId struct {
	Id       string `json:"id,omitempty" validate:"required"`
	Title    string `json:"title,omitempty" validate:"required"`
	Text     string `json:"text,omitempty" validate:"required"`
	Date     string `json:"date,omitempty" validate:"required"`
	Favorite string `json:"favorite,omitempty" validate:"required"`
}

func (noteRequestWithId *NoteRequestWithId) ValidateAndParseFromRequest(c *gin.Context) error {
	if err := c.BindJSON(noteRequestWithId); err != nil {
		return err
	}
	if err := validate.Struct(noteRequestWithId); err != nil {
		return err
	}
	if _, err := utils.ParseBoolFromString(noteRequestWithId.Favorite); err != nil {
		return err
	}
	return nil
}

func (noteRequestWithId *NoteRequestWithId) GetModel(username string) (models.NoteWithId, error) {
	id, err := primitive.ObjectIDFromHex(noteRequestWithId.Id)
	if err != nil {
		return models.NoteWithId{}, err
	}

	date, err := utils.ParseDateFromString(noteRequestWithId.Date)
	if err != nil {
		return models.NoteWithId{}, err
	}

	favorite, err := utils.ParseBoolFromString(noteRequestWithId.Favorite)
	if err != nil {
		return models.NoteWithId{}, err
	}

	return models.NoteWithId{
		Id:       id,
		Username: username,
		Title:    noteRequestWithId.Title,
		Text:     noteRequestWithId.Text,
		Date:     date,
		Favorite: favorite,
	}, nil
}
