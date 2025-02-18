package controllers

import (
	"AP1/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func GetAlbum(c *gin.Context, collection *mongo.Collection) {
	albumID, err := primitive.ObjectIDFromHex(c.Param("album_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
		return
	}
	var album models.Album
	err = collection.FindOne(context.Background(), bson.M{"_id": albumID}).Decode(&album)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, album)
}
