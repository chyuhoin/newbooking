package service

import (
	"newbooking/pkg/dao"
	"newbooking/pkg/entity"
)

type HotelService struct {
	hotelMapper *dao.HotelMapper
}

func NewHotelService() *HotelService {
	return &HotelService{dao.NewHotelMapper()}
}

func (service *HotelService) SearchId(id int) (*entity.Hotel, error) {
	hotel, err := service.hotelMapper.GetHotelById(id)
	if err != nil {
		return nil, err
	}
	return hotel, nil
}

func (service *HotelService) SearchFuzzy(hotel *entity.Hotel) (*[]*entity.Hotel, error) {
	hotels, err := service.hotelMapper.GetHotelsByHotelFuzzy(hotel)
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (service *HotelService) SearchRoom(in *string, out *string, name *string, dest *string, city *string, province *string) (*[]*entity.HotelWithRoom, error) {
	roomList := make([]*entity.HotelWithRoom, 0)
	rooms, err := service.hotelMapper.GetHotelRoom(in, out, name, dest, city, province)
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

func (service *HotelService) SearchRoomAndHotelByRoomId(roomId string) (entity.Room, entity.Hotel, error) {
	room, err := service.hotelMapper.GetRoomByRoomId(roomId)
	if err != nil {
		return entity.Room{}, entity.Hotel{}, err
	}
	hotel, err := service.hotelMapper.GetHotelByHotelId(room.HotelId)
	if err != nil {
		return entity.Room{}, entity.Hotel{}, err
	}
	return room, hotel, nil
}
