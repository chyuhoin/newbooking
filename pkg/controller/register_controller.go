package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/entity"
	"newbooking/pkg/service"
	"newbooking/pkg/utils"
)

type RegisterController struct {
	registerService *service.RegisterService
}

func NewRegisterController() *RegisterController {
	return &RegisterController{registerService: service.NewRegisterService()}
}

func (ctl *RegisterController) GetBooking(c *gin.Context) {
	j := utils.NewJWT()

	token := c.Request.Header.Get("token")
	payload, err := j.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	register, err := ctl.registerService.GetMyRegister(payload.Id)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "books": register})
}

func (ctl *RegisterController) PostBooking(c *gin.Context) {
	var register entity.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "wrong body", "info": err.Error()})
		return
	}

	err := ctl.registerService.AddOneRegister(&register)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}
