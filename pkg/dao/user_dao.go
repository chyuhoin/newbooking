package dao

import (
	"fmt"
	"newbooking/pkg/entity"
	"newbooking/pkg/utils"
	"xorm.io/xorm"
)

type UserMapper struct {
	engine *xorm.Engine
}

func NewUserMapper() *UserMapper {
	return &UserMapper{engine: utils.GetOrmEngine()}
}

func (mapper *UserMapper) GetUserById(id string) (*entity.User, error) {
	var user = entity.User{}
	_, err := mapper.engine.Where("id = ?", id).Get(&user)
	return &user, err
}

func (mapper *UserMapper) GetUserByNameAndPassword(username string, password string) (*entity.User, error) {
	var user = entity.User{}
	_, err := mapper.engine.Where("username = ?", username).And("password = ?", password).Get(&user)
	return &user, err
}

func (mapper *UserMapper) CheckUserWithNameAndPassword(username string, password string) bool {
	has, err := mapper.engine.Where("username = ?", username).And("password = ?", password).Exist(&entity.User{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return has
}

func (mapper *UserMapper) GetAllUsers() ([]*entity.User, error) {
	userList := make([]*entity.User, 0)
	err := mapper.engine.Find(&userList)
	return userList, err
}

func (mapper *UserMapper) InsertUser(user *entity.User) (bool, error) {
	sql := "INSERT INTO t_user (username, password, role) VALUES (?, ?, ?)"
	res, err := mapper.engine.Exec(sql, user.Username, user.Password, user.Role)
	if err != nil {
		return false, err
	}
	rows, err2 := res.RowsAffected()
	return rows > 0, err2
}

func (mapper *UserMapper) UpdateUser(user *entity.User) error {
	newUser := entity.User{
		Email:    user.Email,
		Phone:    user.Phone,
		Gender:   user.Gender,
		Address:  user.Address,
		Nickname: user.Nickname,
	}
	_, err := mapper.engine.ID(user.Id).Update(newUser)
	return err
}

func (mapper *UserMapper) UpdateUserPassword(user *entity.User) error {
	newUser := entity.User{
		Password: user.Password,
	}
	_, err := mapper.engine.ID(user.Id).Update(newUser)
	return err
}
