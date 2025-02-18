package controllers

import (
	"AP1/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"path/filepath"
	"time"
)

type TrackMetadata struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

func GetTracks(c *gin.Context, collection *mongo.Collection) {
	var tracks []models.Track
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tracks"})
		return
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tracks"})
		}
	}(cursor, context.Background())
	for cursor.Next(context.Background()) {
		var track models.Track
		if err := cursor.Decode(&track); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding track"})
			return
		}
		tracks = append(tracks, track)
	}
	c.JSON(http.StatusOK, tracks)
}

func UploadTrack(c *gin.Context, collection *mongo.Collection) {
	// Получаем файл
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить файл"})
		return
	}

	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), fileExt)
	savePath := filepath.Join("uploads", fileName)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	// Получаем JSON-метаданные
	metadataJSON := c.PostForm("metadata")
	var metadata models.Track
	if err := json.Unmarshal([]byte(metadataJSON), &metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат метаданных"})
		return
	}

	// Создаем объект трека
	track := models.Track{
		ID:       primitive.NewObjectID(),
		Title:    metadata.Title,
		Artist:   metadata.Artist,
		Duration: 0, // TODO: Можно дописать обработку длительности через ffmpeg
		Path:     savePath,
		FileName: fileName,
	}

	// Сохраняем в БД
	_, err = collection.InsertOne(context.Background(), track)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении в БД"})
		return
	}

	// Отправляем ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "Файл загружен",
		"track":   track,
	})
}

const UploadDir = "./uploads/"

func GetTrackFile(c *gin.Context) {
	filename := c.Param("filename")

	filePath := filepath.Join(UploadDir, filename)

	c.File(filePath)
}

func GetTrack(c *gin.Context, collection *mongo.Collection) {
	trackID, err := primitive.ObjectIDFromHex(c.Param("track_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid track ID"})
		return
	}
	var track models.Track
	err = collection.FindOne(context.Background(), bson.M{"_id": trackID}).Decode(&track)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}
	c.JSON(http.StatusOK, track)
}

func DeleteTrack(c *gin.Context, collection *mongo.Collection) {
	trackID, _ := primitive.ObjectIDFromHex(c.Param("track_id"))
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": trackID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete track"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Track deleted successfully"})
}
