package entity

type Room struct {
	Id       string  `json:"id"`
	HotelId  int     `json:"hotelId"`
	Name     string  `json:"name"`
	Bed      string  `json:"bed"`
	Capacity string  `json:"capacity"`
	Price    float64 `json:"price"`
	Num      int     `json:"num"`
	Remain   int     `json:"remain" xorm:"-"`
}
