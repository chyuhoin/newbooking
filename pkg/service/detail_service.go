package service

import (
	"newbooking/pkg/dao"
	"newbooking/pkg/entity"
)

type DetailService struct {
	detailMapper *dao.DetailMapper
}

func NewDetailService() *DetailService {
	return &DetailService{detailMapper: dao.NewDetailMapper()}
}

func (service *DetailService) ViewImages(hotelId int) (*entity.ImageGroup, error) {
	return service.detailMapper.GetImagesByHotelId(hotelId)
}
