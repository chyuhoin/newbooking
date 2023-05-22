package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/service"
	"strconv"
)

type AdminController struct {
	hotelService    *service.HotelService
	registerService *service.RegisterService
	userService     *service.UserService
}

func NewAdminController() *AdminController {
	return &AdminController{
		hotelService:    service.NewHotelService(),
		registerService: service.NewRegisterService(),
		userService:     service.NewUserService(),
	}
}

func (ctl *AdminController) Hotel(c *gin.Context) {
	dest := c.Query("dest")

	hotels, err := ctl.hotelService.SearchHotelByDest(dest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "hotels": hotels})
}

func (ctl *AdminController) DeleteHotel(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id is not number", "info": err.Error()})
		return
	}

	err = ctl.hotelService.RemoveOneHotel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (ctl *AdminController) HotelRegister(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "id is not number", "info": err.Error()})
		return
	}

	regs, err := ctl.registerService.ListRegisterByHotel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "register": regs})
}

func (ctl *AdminController) DeleteUser(c *gin.Context) {
	id := c.Query("id")

	err := ctl.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
