package dao

import (
	"newbooking/pkg/entity"
	"newbooking/pkg/utils"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type RegisterMapper struct {
	engine *xorm.Engine
}

func NewRegisterMapper() *RegisterMapper {
	return &RegisterMapper{engine: utils.GetOrmEngine()}
}

func (mapper *RegisterMapper) GetRegisterByUserId(userId string) (*[]*entity.Register, error) {
	registerInfo := make([]*entity.Register, 0)
	err := mapper.engine.Where(builder.Eq{"user_id": userId}).Find(&registerInfo)
	return &registerInfo, err
}
