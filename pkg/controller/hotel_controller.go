package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbooking/pkg/entity"
	"newbooking/pkg/service"
	"strconv"
)

type HotelController struct {
	hotelService *service.HotelService
}

func NewHotelController() *HotelController {
	return &HotelController{hotelService: service.NewHotelService()}
}

func (ctl *HotelController) List(c *gin.Context) {
	var hotel = entity.Hotel{}
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	res := ctl.hotelService.SearchFuzzy(&hotel)
	c.JSON(http.StatusOK, gin.H{"msg": "success", "hotels": res})
}

func (ctl *HotelController) Search(c *gin.Context) {
	dest := c.Query("dest")
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")
	adult, _ := strconv.Atoi(c.Query("adult"))
	children, _ := strconv.Atoi(c.Query("children"))

	hotels := make([]map[string]interface{}, 0)
	rooms, err := ctl.hotelService.SearchRoom(&checkin, &checkout, &dest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	for _, room := range *rooms {
		cost := room.CalcMinCost(adult, children)
		hotel := make(map[string]interface{})
		hotel["info"] = room.Hotel
		hotel["cost"] = cost
		hotels = append(hotels, hotel)
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "hotels": hotels})
}
