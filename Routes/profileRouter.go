package routes

import (
	controller "mora-Backend/controllers"
	"mora-Backend/middleware"
	
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/profiles/signup", controller.SignUp())
	incomingRoutes.POST("/profiles/login", controller.Login())
	incomingRoutes.POST("/profiles/logout", middleware.AuthMiddleware(), controller.Logout())
	incomingRoutes.GET("/profiles/", middleware.AuthMiddleware(), controller.GetProfiles())
}