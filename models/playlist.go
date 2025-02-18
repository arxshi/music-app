package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	User   primitive.ObjectID   `bson:"user"`
	Name   string               `bson:"name"`
	Tracks []primitive.ObjectID `bson:"tracks"`
}
