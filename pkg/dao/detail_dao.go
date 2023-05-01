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

func (mapper *DetailMapper) GetImagesByHotelId(hotelId int) (*entity.ImageGroup, error) {
	imageList := make([]string, 0)
	err := mapper.engine.Table("t_image").Where(builder.Eq{"hotel_id": hotelId}).Cols("image").Find(&imageList)
	if err != nil {
		return &entity.ImageGroup{}, err
	}
	return &entity.ImageGroup{
		HotelId: hotelId,
		Images:  &imageList,
	}, nil
}
