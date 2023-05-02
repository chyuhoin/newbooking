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
		return
	}
	images, err := ctl.service.ViewImages(hotelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "database crashed", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "images": images.Images})
}

func (ctl *DetailController) Description(c *gin.Context) {
	hotelId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong number"})
		return
	}
	description, err := ctl.service.GetDescription(hotelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "database crashed", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "description": description})
}

func (ctl *DetailController) Policy(c *gin.Context) {
	hotelId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong number"})
		return
	}
	policy, err := ctl.service.GetPolicy(hotelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "database crashed", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "policy": policy})
}

func (ctl *DetailController) Notes(c *gin.Context) {
	hotelId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Wrong number"})
		return
	}
	notes, err := ctl.service.GetNotes(hotelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "database crashed", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "notes": notes})
}
