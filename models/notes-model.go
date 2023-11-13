package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Username string
	Title    string
	Text     string
	Date     time.Time
	Favorite bool
}

type NoteWithId struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string
	Title    string
	Text     string
	Date     time.Time
	Favorite bool
}
