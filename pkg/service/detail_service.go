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

func (service *DetailService) GetDescription(hotelId int) (*string, error) {
	return service.detailMapper.GetDescriptionByHotelId(hotelId)
}

func (service *DetailService) ViewImages(hotelId int) (*entity.ImageGroup, error) {
	return service.detailMapper.GetImagesByHotelId(hotelId)
}

func (service *DetailService) GetPolicy(hotelId int) (*entity.Policy, error) {
	return service.detailMapper.GetPolicyByHotelId(hotelId)
}

func (service *DetailService) GetNotes(hotelId int) (*entity.Notes, error) {
	return service.detailMapper.GetNotesByHotelId(hotelId)
}
