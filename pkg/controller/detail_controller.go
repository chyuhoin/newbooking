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
	images, err := ctl.service.ViewImages(hotelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "database crashed", "info": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "success", "images": images.Images})
	}
}
