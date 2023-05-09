package dao

import (
	"newbooking/pkg/entity"
	"newbooking/pkg/utils"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type DetailMapper struct {
	engine *xorm.Engine
}

func NewDetailMapper() *DetailMapper {
	return &DetailMapper{engine: utils.GetOrmEngine()}
}

func (mapper *DetailMapper) GetDescriptionByHotelId(hotelId int) (*string, error) {
	var description string
	_, err := mapper.engine.Table("t_desc").
		Where(builder.Eq{"id": hotelId}).Cols("description").Get(&description)
	return &description, err
}

func (mapper *DetailMapper) GetImagesByHotelId(hotelId int) (*entity.ImageGroup, error) {
	imageList := make([]string, 0)
	err := mapper.engine.Table("t_image").
		Where(builder.Eq{"hotel_id": hotelId}).Cols("image").Find(&imageList)
	return &entity.ImageGroup{
		HotelId: hotelId,
		Images:  &imageList,
	}, err
}

func (mapper *DetailMapper) GetPolicyByHotelId(hotelId int) (*entity.Policy, error) {
	var policy entity.Policy
	_, err := mapper.engine.Table("t_policy").Where(builder.Eq{"id": hotelId}).Get(&policy)
	return &policy, err
}

func (mapper *DetailMapper) GetNotesByHotelId(hotelId int) (*entity.Notes, error) {
	var notes entity.Notes
	_, err := mapper.engine.Table("t_note").Where(builder.Eq{"id": hotelId}).Get(&notes)
	return &notes, err
}

func (mapper *DetailMapper) GetRoomsByHotelId(hotelId int) (*[]*entity.Room, error) {
	rooms := make([]*entity.Room, 0)
	err := mapper.engine.Table("t_room").
		Where(builder.Eq{"hotel_id": hotelId}).
		Cols("id", "hotel_id", "name", "bed", "capacity", "price", "num").
		Find(&rooms)
	return &rooms, err
}
