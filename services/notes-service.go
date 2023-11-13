package services

import (
	"context"
	"notes-appapi/config"
	"notes-appapi/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var notesCollection = config.GetCollection(config.DB, "notes")

func GetUsersNotes(ctx context.Context, username string) ([]models.NoteWithId, error) {
	var notes []models.NoteWithId
	cursor, err := notesCollection.Find(ctx, bson.M{"username": username})
	if err != nil {
		return notes, err
	}
	if err := cursor.All(ctx, &notes); err != nil {
		return notes, err
	}
	return notes, nil
}

func GetUsersRecentNotes(ctx context.Context, username string) ([]models.NoteWithId, error) {
	var notes []models.NoteWithId
	opts := options.Find().SetSort(bson.M{"date": -1})
	cursor, err := notesCollection.Find(ctx, bson.M{"username": username}, opts)
	if err != nil {
		return notes, err
	}
	if err := cursor.All(ctx, &notes); err != nil {
		return notes, err
	}
	return notes, nil
}

func GetUsersFavoriteNotes(ctx context.Context, username string) ([]models.NoteWithId, error) {
	var notes []models.NoteWithId
	opts := options.Find().SetSort(bson.M{"date": -1})
	cursor, err := notesCollection.Find(ctx, bson.M{"username": username, "favorite": true}, opts)
	if err != nil {
		return notes, err
	}
	if err := cursor.All(ctx, &notes); err != nil {
		return notes, err
	}
	return notes, nil
}

func SetNoteFavorite(ctx context.Context, noteId string, favorite bool) error {
	id, err := primitive.ObjectIDFromHex(noteId)
	if err != nil {
		return err
	}

	update := bson.D{{"$set", bson.D{{"favorite", favorite}}}}
	result, err := notesCollection.UpdateByID(ctx, id, update)
	println(result)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNote(ctx context.Context, note models.NoteWithId) error {
	filter := bson.M{"_id": note.Id}
	update := bson.M{"$set": note}
	result, err := notesCollection.UpdateOne(ctx, filter, update)
	println(result)
	if err != nil {
		return err
	}
	return nil
}
