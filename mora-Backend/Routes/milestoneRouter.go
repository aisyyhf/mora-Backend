package routes

import (
	controller "golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

func MilestoneRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/milestones", middleware.AuthMiddleware(), controller.GetMilestones())
	incomingRoutes.GET("/milestones/:milestone_id", middleware.AuthMiddleware(), controller.GetMilestone())
	incomingRoutes.POST("/milestones", middleware.AuthMiddleware(), controller.CreateMilestone())
	incomingRoutes.PUT("/milestones/:milestone_id", middleware.AuthMiddleware(), controller.UpdateMilestone())
	incomingRoutes.DELETE("/milestones/:milestone_id", middleware.AuthMiddleware(), controller.DeleteMilestone())
}
