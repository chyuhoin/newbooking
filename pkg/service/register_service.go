package service

import (
	"newbooking/pkg/dao"
	"newbooking/pkg/entity"
)

type RegisterService struct {
	registerMapper *dao.RegisterMapper
}

func NewRegisterService() *RegisterService {
	return &RegisterService{registerMapper: dao.NewRegisterMapper()}
}

func (service *RegisterService) GetMyRegister(userId string) (*[]*entity.Register, error) {
	return service.registerMapper.GetRegisterByUserId(userId)
}

func (service *RegisterService) AddOneRegister(register *entity.Register) error {
	return service.registerMapper.InsertRegister(register)
}

func (service *RegisterService) CancelRegister(register *entity.Register) error {
	return service.registerMapper.RemoveRegister(register)
}

func (service *RegisterService) ModifyRegisterTime(register *entity.Register) error {
	return service.registerMapper.UpdateRegisterTime(register)
}
