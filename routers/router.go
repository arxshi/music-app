package routers

import (
	"AP1/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(router *gin.Engine, db *mongo.Database) {
	router.POST("/auth/register", func(c *gin.Context) { controllers.RegisterUser(c, db.Collection("users")) })
	router.POST("/auth/login", func(c *gin.Context) { controllers.Login(c, db.Collection("users")) })
	router.GET("/users/:user_id", func(c *gin.Context) { controllers.GetUser(c, db.Collection("users")) })
	router.GET("/users", func(c *gin.Context) { controllers.GetUser(c, db.Collection("users")) })
	router.GET("/tracks", func(c *gin.Context) { controllers.GetTracks(c, db.Collection("tracks")) })
	router.GET("/tracks/:track_id", func(c *gin.Context) { controllers.GetTrack(c, db.Collection("tracks")) })
	router.DELETE("/tracks/:track_id", func(c *gin.Context) { controllers.DeleteTrack(c, db.Collection("tracks")) })
	router.POST("/tracks", func(c *gin.Context) { controllers.UploadTrack(c, db.Collection("tracks")) })
	router.GET("/albums/:album_id", func(c *gin.Context) { controllers.GetAlbum(c, db.Collection("albums")) })
	router.GET("/playlists/:playlist_id", func(c *gin.Context) { controllers.GetPlaylistByID(c, db.Collection("playlists")) })
	router.POST("/playlists", func(c *gin.Context) { controllers.CreatePlaylist(c, db.Collection("playlists")) })
	router.DELETE("/playlists/:playlist_id/tracks/:track_id", func(c *gin.Context) { controllers.RemoveTrackFromPlaylist(c, db.Collection("playlists")) })
	router.POST("/playlists/:playlist_id/tracks/:track_id", func(c *gin.Context) { controllers.AddTrackToPlaylist(c, db.Collection("playlists")) })
	router.POST("/reviews", func(c *gin.Context) { controllers.CreateReview(c, db.Collection("reviews")) })
	router.GET("/reviews", func(c *gin.Context) { controllers.GetReviews(c, db.Collection("reviews")) })
	router.POST("/upload", func(c *gin.Context) { controllers.UploadTrack(c, db.Collection("tracks")) })
	router.GET("/uploads/:filename", controllers.GetTrackFile)
}
