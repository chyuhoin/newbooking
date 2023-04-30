package test

import (
	"fmt"
	"newbooking/pkg/service"
	"testing"
	"time"
)

func TestService(t *testing.T) {
	hotelService := service.NewHotelService()
	in, _ := time.Parse("2006-01-02", "2023-01-02")
	out, _ := time.Parse("2006-01-02", "2023-03-02")
	city := "南京"
	rooms := hotelService.SearchRoom(&in, &out, &city)
	for _, room := range *rooms {
		fmt.Printf("%d %f\n", room.Id, room.CalcMinCost(3, 2))
	}
}
