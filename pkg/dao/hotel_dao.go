package dao

import (
	"fmt"
	"newbooking/pkg/entity"
	"newbooking/pkg/utils"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type RoomTmp struct {
	entity.Room  `xorm:"extends"`
	entity.Hotel `xorm:"extends"`
}

type HotelMapper struct {
	engine *xorm.Engine
}

func NewHotelMapper() *HotelMapper {
	return &HotelMapper{engine: utils.GetOrmEngine()}
}

func (mapper *HotelMapper) GetHotelById(id int) (*entity.Hotel, error) {
	var hotel entity.Hotel
	_, err := mapper.engine.Where("id = ?", id).Get(&hotel)
	return &hotel, err
}

func (mapper *HotelMapper) GetHotelsByHotelFuzzy(hotel *entity.Hotel) (*[]*entity.Hotel, error) {
	hotelList := make([]*entity.Hotel, 0)
	err := mapper.engine.
		Where(builder.Like{"province", hotel.Province}).
		And(builder.Like{"name", hotel.Name}).
		And(builder.Like{"location", hotel.Location}).
		And(builder.Like{"city", hotel.City}).Find(&hotelList)
	return &hotelList, err
}

func (mapper *HotelMapper) GetHotelRoom(in *string, out *string, dest *string, city *string, province *string) (*[]*RoomTmp, error) {
	hotelList := make([]*RoomTmp, 0)
	SQL := `
	SELECT * FROM (
		SELECT
			rom.*,
			num - COALESCE ( regcnt, 0 ) remain
		FROM (
			SELECT COUNT
				( * ) regcnt,
				room_id rid
			FROM
				t_register
			WHERE
				( NOT ( end_time < '%s' OR start_time > '%s' ) )
				AND is_deleted = FALSE
			GROUP BY
				room_id
			) reg
		RIGHT JOIN (
			SELECT
				*
			FROM
				t_room
			WHERE
				hotel_id IN (
					SELECT
						"id"
					FROM
						t_hotel
					WHERE
						((name LIKE '%s' OR location LIKE '%s') AND city LIKE '%s' AND province LIKE '%s')
				)
			) rom 
		ON
			rom."id" = rid
		WHERE
			num - COALESCE ( regcnt, 0 ) != 0
		) TMP
	INNER JOIN
		t_hotel
	ON
		TMP.hotel_id = t_hotel.id;
	`
	SQL = fmt.Sprintf(SQL, *in, *out, "%"+*dest+"%", "%"+*dest+"%", "%"+*city+"%", "%"+*province+"%")
	err := mapper.engine.SQL(SQL).Find(&hotelList)
	return &hotelList, err
}

func (mapper *HotelMapper) GetRoomByRoomId(roomId string) (entity.Room, error) {
	var room entity.Room
	_, err := mapper.engine.
		Cols("id", "hotel_id", "name", "bed", "capacity", "price", "num").
		Where(builder.Eq{"id": roomId}).Get(&room)
	return room, err
}

func (mapper *HotelMapper) GetHotelByHotelId(hotelId int) (entity.Hotel, error) {
	var hotel entity.Hotel
	_, err := mapper.engine.Where(builder.Eq{"id": hotelId}).Get(&hotel)
	return hotel, err
}
