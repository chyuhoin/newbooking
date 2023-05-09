package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/entity"
	"newbooking/pkg/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	controller := UserController{service.NewUserService()}
	return &controller
}

func (ctl *UserController) Login(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if token, err := ctl.userService.Login(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "No such people"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "success", "token": token, "user": user})
	}
}

func (ctl *UserController) Register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if ok, err := ctl.userService.RegisterByNameAndPassword(&user.Username, &user.Password); ok {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}
}

func (ctl *UserController) Users(c *gin.Context) {
	list := ctl.userService.ListUsers()
	c.JSON(http.StatusOK, gin.H{"users": list})
}
