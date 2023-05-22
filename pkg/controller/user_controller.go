package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/entity"
	"newbooking/pkg/service"
	"newbooking/pkg/utils"
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

func (ctl *UserController) UserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	tokenInfo := claims.(*utils.CustomClaims)
	user, err := ctl.userService.GetOneUserInfo(tokenInfo.Id)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "info": user})
}

func (ctl *UserController) PutUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	claims, _ := c.Get("claims")
	tokenInfo := claims.(*utils.CustomClaims)
	user.Id = tokenInfo.Id

	if err := ctl.userService.UpdateUserInfo(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "something wrong", "info": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
}

func (ctl *UserController) ChangePassword(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	claims, _ := c.Get("claims")
	tokenInfo := claims.(*utils.CustomClaims)
	user.Id = tokenInfo.Id

	if err := ctl.userService.UpdateUserPassword(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "something wrong", "info": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	}
}
