package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/service"
	"strconv"
)

type DetailController struct {
	service *service.DetailService
}

func NewDetailController() *DetailController {
	return &DetailController{service: service.NewDetailService()}
}

func (ctl *DetailController) Images(c *gin.Context) {
	hotelId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong number"})
	}
	images := ctl.service.ViewImages(hotelId)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "images": images.Images})
}
