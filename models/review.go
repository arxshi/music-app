package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	User   primitive.ObjectID `bson:"user"`
	Track  primitive.ObjectID `bson:"track"`
	Review string             `bson:"review"`
	Rating int                `bson:"rating"`
}
