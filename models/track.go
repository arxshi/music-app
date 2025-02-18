package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Track struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Artist   string             `bson:"artist"`
	AlbumID  primitive.ObjectID `bson:"album,omitempty"`
	Duration int                `bson:"duration"`
	Path     string             `bson:"path"`
	FileName string             `bson:"filename"`
}
