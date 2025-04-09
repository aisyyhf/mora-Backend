package routes

import (
	controller "mora-Backend/controllers"
	"mora-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func MemoryRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/memories", middleware.AuthMiddleware(), controller.CreateMemory())
	incomingRoutes.GET("/memories", middleware.AuthMiddleware(), controller.GetMemories())
	incomingRoutes.GET("/memories/:memory_id", middleware.AuthMiddleware(), controller.GetMemoryByID())
	incomingRoutes.PUT("/memories/:memory_id", middleware.AuthMiddleware(), controller.UpdateMemory())
	incomingRoutes.DELETE("/memories/:memory_id", middleware.AuthMiddleware(), controller.DeleteMemory())
}
