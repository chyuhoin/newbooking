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
	var hotel = entity.Hotel{
		Name:     c.Query("name"),
		Province: c.Query("province"),
		City:     c.Query("city"),
		Location: c.Query("location"),
	}

	res, err := ctl.hotelService.SearchFuzzy(&hotel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "hotels": res})
}

func (ctl *HotelController) Search(c *gin.Context) {
	dest := c.Query("dest")
	city := c.Query("city")
	province := c.Query("province")
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")
	adult, _ := strconv.Atoi(c.Query("adult"))
	children, _ := strconv.Atoi(c.Query("children"))

	hotels := make([]map[string]interface{}, 0)
	rooms, err := ctl.hotelService.SearchRoom(&checkin, &checkout, &dest, &city, &province)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}

	for _, room := range *rooms {
		cost, method := room.CalcMinCost(adult, children)
		hotel := make(map[string]interface{})
		hotel["info"] = room.Hotel
		hotel["cost"] = cost
		hotel["method"] = method
		hotels = append(hotels, hotel)
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "hotels": hotels})
}

func (ctl *HotelController) Room(c *gin.Context) {
	id := c.Query("id")
	room, hotel, err := ctl.hotelService.SearchRoomAndHotelByRoomId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "something wrong", "info": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success", "room": room, "hotel": hotel})
}
