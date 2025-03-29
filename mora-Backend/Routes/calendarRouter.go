package routes

import (
	controller "golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

func CalendarRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/calendar", middleware.AuthMiddleware(), controller.GetCalendarEvents())
	incomingRoutes.GET("/calendar/:event_id", middleware.AuthMiddleware(), controller.GetCalendarEvent())
	incomingRoutes.POST("/calendar", middleware.AuthMiddleware(), controller.CreateCalendarEvent())
	incomingRoutes.PUT("/calendar/:event_id", middleware.AuthMiddleware(), controller.UpdateCalendarEvent())
	incomingRoutes.DELETE("/calendar/:event_id", middleware.AuthMiddleware(), controller.DeleteCalendarEvent())
}
