package routes

import (
	controller "golang-restaurant-management/controllers"
	"golang-restaurant-management/middleware"

	"github.com/gin-gonic/gin"
)

func WishlistRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/wishlists", middleware.AuthMiddleware(), controller.GetWishlists())          // Mengambil semua wishlist pasangan
	incomingRoutes.GET("/wishlists/:wishlist_id", middleware.AuthMiddleware(), controller.GetWishlist()) // Mengambil detail wishlist berdasarkan ID
	incomingRoutes.POST("/wishlists", middleware.AuthMiddleware(), controller.CreateWishlist())        // Membuat wishlist baru untuk pasangan
	incomingRoutes.PUT("/wishlists/:wishlist_id", middleware.AuthMiddleware(), controller.UpdateWishlist()) // Mengedit wishlist berdasarkan ID
	incomingRoutes.DELETE("/wishlists/:wishlist_id", middleware.AuthMiddleware(), controller.DeleteWishlist()) // Menghapus wishlist berdasarkan ID
}
