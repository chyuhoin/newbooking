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

	router.POST("/login", userController.Login)
	router.GET("/users", userController.Users)

	hotel := router.Group("/hotel")
	hotel.GET("/list", hotelController.List)
	hotel.GET("/search", hotelController.Search)

}
