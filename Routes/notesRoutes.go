package routes

import (
	controller "mora-Backend/controllers"
	"mora-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func NotesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", middleware.AuthMiddleware(), controller.GetNotes())
	incomingRoutes.GET("/notes/:note_id", middleware.AuthMiddleware(), controller.GetNote())
	incomingRoutes.POST("/notes", middleware.AuthMiddleware(), controller.CreateNote())
	incomingRoutes.PUT("/notes/:note_id", middleware.AuthMiddleware(), controller.UpdateNote())
	incomingRoutes.DELETE("/notes/:note_id", middleware.AuthMiddleware(), controller.DeleteNote())
}
