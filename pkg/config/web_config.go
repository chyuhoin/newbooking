package config

import (
	"github.com/gin-gonic/gin"
	"newbooking/pkg/controller"
	"newbooking/pkg/middleware"
)

func RouterConfig(router *gin.Engine) {
	router.Use(middleware.JWTAuth())
	userController := controller.NewUserController()
	hotelController := controller.NewHotelController()
	detailController := controller.NewDetailController()

	router.POST("/login", userController.Login)
	router.GET("/users", userController.Users)

	hotel := router.Group("/hotel")
	hotel.GET("/list", hotelController.List)
	hotel.GET("/search", hotelController.Search)

	detail := router.Group("/detail")
	detail.GET("/images", detailController.Images)
	detail.GET("/description", detailController.Description)
	detail.GET("/policy", detailController.Policy)
	detail.GET("/notes", detailController.Notes)

}
