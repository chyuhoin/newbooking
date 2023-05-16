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

func (mapper *RegisterMapper) InsertRegister(register *entity.Register) error {
	SQL := `
INSERT INTO "t_register" 
("room_id","start_time","end_time","user_id","b_name","b_email","b_state","b_phone"
,"b_birthday","r_name","r_email","plan") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`
	_, err := mapper.engine.Exec(SQL, register.RoomId, register.StartTime, register.EndTime,
		register.UserId, register.BookerName, register.BookerEmail, register.BookerState,
		register.BookerPhone, register.BookerBirthday, register.RoomerName, register.RoomerEmail,
		register.Plan)
	return err
}

func (mapper *RegisterMapper) RemoveRegister(register *entity.Register) error {
	_, err := mapper.engine.ID(register.Id).Cols("is_deleted").Update(register)
	return err
}

func (mapper *RegisterMapper) UpdateRegisterTime(register *entity.Register) error {
	_, err := mapper.engine.ID(register.Id).Cols("start_time", "end_time").Update(register)
	return err
}
