package routes

import (
	controller "mora-Backend/controllers"
	"mora-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func WishlistRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/wishlists", middleware.AuthMiddleware(), controller.GetWishlists())          
	incomingRoutes.GET("/wishlists/:wishlist_id", middleware.AuthMiddleware(), controller.GetWishlist()) 
	incomingRoutes.POST("/wishlists", middleware.AuthMiddleware(), controller.CreateWishlist())        
	incomingRoutes.PUT("/wishlists/:wishlist_id", middleware.AuthMiddleware(), controller.UpdateWishlist()) 
	incomingRoutes.DELETE("/wishlists/:wishlist_id", middleware.AuthMiddleware(), controller.DeleteWishlist())
}
