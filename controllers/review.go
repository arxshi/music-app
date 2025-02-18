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

func CreateReview(c *gin.Context, collection *mongo.Collection) {
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	review.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}
	c.JSON(http.StatusOK, review)
}

func GetReviews(c *gin.Context, collection *mongo.Collection) {
	var reviews []models.Review
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reviews"})
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var review models.Review
		if err := cursor.Decode(&review); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding review"})
			return
		}
		reviews = append(reviews, review)
	}
	c.JSON(http.StatusOK, reviews)
}
