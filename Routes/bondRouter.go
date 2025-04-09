package routes

import (
	controller "mora-Backend/controllers"
	"mora-Backend/middleware"
	
	"github.com/gin-gonic/gin"
)

func BondRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/bonds/createbonds", middleware.AuthMiddleware(), controller.CreateBond())
	incomingRoutes.POST("/bonds/joinbonds", middleware.AuthMiddleware(), controller.JoinBond())
	incomingRoutes.GET("/bonds/", middleware.AuthMiddleware(), controller.GetBond())
}