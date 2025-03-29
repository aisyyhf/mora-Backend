package routes

import (
	controller "golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"
	
	"github.com/gin-gonic/gin"
)

func BondRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/bonds/createbonds", middleware.AuthMiddleware(), controller.CreateBond())
	incomingRoutes.POST("/bonds/joinbonds", middleware.AuthMiddleware(), controller.JoinBond())
	incomingRoutes.GET("/bonds/", middleware.AuthMiddleware(), controller.GetBond())
}