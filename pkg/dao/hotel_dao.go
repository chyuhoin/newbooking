package dao

import (
	"fmt"
	"newbooking/pkg/entity"
	"newbooking/pkg/utils"
	"time"
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

func (mapper *HotelMapper) GetHotelRoom(in *time.Time, out *time.Time, city *string) (*[]*RoomTmp, error) {
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
						city LIKE'%s'
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
	SQL = fmt.Sprintf(SQL, in.Format("2006-01-02"), out.Format("2006-01-02"), "%"+*city+"%")
	err := mapper.engine.SQL(SQL).Find(&hotelList)
	return &hotelList, err
}