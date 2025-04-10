package routes

import (
	controller "mora-Backend/controllers"
	"mora-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func MilestoneRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/milestones", middleware.AuthMiddleware(), controller.GetMilestones())
	incomingRoutes.GET("/milestones/:milestone_id", middleware.AuthMiddleware(), controller.GetMilestone())
	incomingRoutes.POST("/milestones", middleware.AuthMiddleware(), controller.CreateMilestone())
	incomingRoutes.PUT("/milestones/:milestone_id", middleware.AuthMiddleware(), controller.UpdateMilestone())
	incomingRoutes.DELETE("/milestones/:milestone_id", middleware.AuthMiddleware(), controller.DeleteMilestone())
}
