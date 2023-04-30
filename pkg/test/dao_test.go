package test

import (
	"fmt"
	"newbooking/pkg/dao"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	mapper := dao.NewHotelMapper()
	in, _ := time.Parse("2006-01-02", "2023-01-02")
	out, _ := time.Parse("2006-01-02", "2023-03-02")
	city := "南京"
	hotels, err := mapper.GetHotelRoom(&in, &out, &city)
	if err != nil {
		return
	}
	for i := 0; i < len(*hotels); i++ {
		hotel := (*hotels)[i]
		println("-------------------")
		fmt.Println(hotel)
	}
}
