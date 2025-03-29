package routes

import (
	controller "golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"
	
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/profiles/signup", controller.SignUP())
	incomingRoutes.POST("/profiles/login", controller.Login())
	incomingRoutes.POST("/profiles/logout", middleware.AuthMiddleware(), controller.Logout())
	incomingRoutes.GET("/profiles/", middleware.AuthMiddleware(), controller.GetProfiles())
}