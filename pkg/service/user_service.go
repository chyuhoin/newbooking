package service

import (
	"newbooking/pkg/dao"
	"newbooking/pkg/entity"
	"newbooking/pkg/utils"
)

type UserService struct {
	userMapper *dao.UserMapper
}

func NewUserService() *UserService {
	return &UserService{dao.NewUserMapper()}
}

func (service *UserService) Login(user *entity.User) (*string, error) {
	res, err := service.userMapper.GetUserByNameAndPassword(user.Username, user.Password)
	if res.Id == "" || err != nil {
		return nil, utils.NoSuchUser
	}
	token, err := utils.GenerateToken(res, 60*24)
	if err != nil {
		return nil, err
	}
	user.Id = res.Id
	user.Role = res.Role
	return token, nil
}

func (service *UserService) ListUsers() []*entity.User {
	res, err := service.userMapper.GetAllUsers()
	if err != nil {
		return nil
	}
	return res
}

func (service *UserService) RegisterByNameAndPassword(username *string, password *string) (bool, error) {
	user := entity.User{Username: *username, Password: *password, Role: "user"}
	return service.userMapper.InsertUser(&user)
}

func (service *UserService) GetOneUserInfo(id string) (*entity.User, error) {
	return service.userMapper.GetUserById(id)
}

func (service *UserService) UpdateUserInfo(user *entity.User) error {
	return service.userMapper.UpdateUser(user)
}

func (service *UserService) UpdateUserPassword(user *entity.User) error {
	return service.userMapper.UpdateUserPassword(user)
}

func (service *UserService) DeleteUser(id string) error {
	return service.userMapper.DeleteUserById(id)
}
