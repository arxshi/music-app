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

func GetPlaylistByID(c *gin.Context, collection *mongo.Collection) {
	playlistID, err := primitive.ObjectIDFromHex(c.Param("playlist_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	var playlist models.Playlist
	err = collection.FindOne(context.Background(), bson.M{"_id": playlistID}).Decode(&playlist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Playlist not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve playlist"})
		}
		return
	}

	c.JSON(http.StatusOK, playlist)
}

func GetPlaylists(c *gin.Context, collection *mongo.Collection) {
	userID, _ := primitive.ObjectIDFromHex(c.Param("user_id"))
	var playlists []models.Playlist
	cursor, err := collection.Find(context.Background(), bson.M{"user": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve playlists"})
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var playlist models.Playlist
		if err := cursor.Decode(&playlist); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding playlist"})
			return
		}
		playlists = append(playlists, playlist)
	}
	c.JSON(http.StatusOK, playlists)
}

func CreatePlaylist(c *gin.Context, collection *mongo.Collection) {
	var playlist models.Playlist
	if err := c.BindJSON(&playlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	playlist.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), playlist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create playlist"})
		return
	}
	c.JSON(http.StatusOK, playlist)
}

func RemoveTrackFromPlaylist(c *gin.Context, collection *mongo.Collection) {
	playlistID, _ := primitive.ObjectIDFromHex(c.Param("playlist_id"))
	trackID, _ := primitive.ObjectIDFromHex(c.Param("track_id"))
	update := bson.M{"$pull": bson.M{"tracks": trackID}}
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": playlistID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove track from playlist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Track removed from playlist"})
}

func AddTrackToPlaylist(c *gin.Context, collection *mongo.Collection) {
	playlistID, _ := primitive.ObjectIDFromHex(c.Param("playlist_id"))
	trackID, _ := primitive.ObjectIDFromHex(c.Param("track_id"))
	update := bson.M{"$addToSet": bson.M{"tracks": trackID}}
	_, err := collection.UpdateOne(context.Background(), bson.M{"_id": playlistID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add track to playlist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Track added to playlist"})
}
