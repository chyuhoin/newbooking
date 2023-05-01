package service

import (
	"newbooking/pkg/dao"
	"newbooking/pkg/entity"
)

type HotelService struct {
	userMapper *dao.HotelMapper
}

func NewHotelService() *HotelService {
	return &HotelService{dao.NewHotelMapper()}
}

func (service *HotelService) SearchId(id int) *entity.Hotel {
	hotel, err := service.userMapper.GetHotelById(id)
	if err != nil {
		return nil
	}
	return hotel
}

func (service *HotelService) SearchFuzzy(hotel *entity.Hotel) *[]*entity.Hotel {
	hotels, err := service.userMapper.GetHotelsByHotelFuzzy(hotel)
	if err != nil {
		return nil
	}
	return hotels
}

func (service *HotelService) SearchRoom(in *string, out *string, city *string) (*[]*entity.HotelWithRoom, error) {
	roomList := make([]*entity.HotelWithRoom, 0)
	rooms, err := service.userMapper.GetHotelRoom(in, out, city)
	if err != nil {
		return nil, err
	}
	for _, room := range *rooms {
		tmp := room.Room
		if len(roomList) != 0 && room.HotelId == roomList[len(roomList)-1].Id {
			roomList[len(roomList)-1].Rooms = append(roomList[len(roomList)-1].Rooms, tmp)
		} else {
			roomList = append(roomList, &entity.HotelWithRoom{
				Hotel: room.Hotel,
				Rooms: []entity.Room{tmp},
			})
		}
	}
	return &roomList, nil
}
