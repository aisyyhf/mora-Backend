package routes

import (
	"golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

func MemoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/memories", middleware.AuthMiddleware(), controllers.CreateMemory())
	incomingRoutes.GET("/memories", middleware.AuthMiddleware(), controllers.GetMemories())
	incomingRoutes.GET("/memories/:memory_id", middleware.AuthMiddleware(), controllers.GetMemoryByID())
	incomingRoutes.PUT("/memories/:memory_id", middleware.AuthMiddleware(), controllers.UpdateMemory())
	incomingRoutes.DELETE("/memories/:memory_id", middleware.AuthMiddleware(), controllers.DeleteMemory())
}
