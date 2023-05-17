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
	registerController := controller.NewRegisterController()

	router.POST("/login", userController.Login)
	router.POST("/register", userController.Register)
	router.GET("/users", userController.Users)

	hotel := router.Group("/hotel")
	hotel.GET("/list", hotelController.List)
	hotel.GET("/search", hotelController.Search)
	hotel.GET("/room", hotelController.Room)

	detail := router.Group("/detail")
	detail.GET("/images", detailController.Images)
	detail.GET("/description", detailController.Description)
	detail.GET("/policy", detailController.Policy)
	detail.GET("/notes", detailController.Notes)
	detail.GET("/room", detailController.Room)

	book := router.Group("/book")
	book.GET("", registerController.GetBooking)
	book.POST("", registerController.PostBooking)
	book.DELETE("", registerController.DeleteBooking)
	book.PUT("", registerController.PutBooking)

}
