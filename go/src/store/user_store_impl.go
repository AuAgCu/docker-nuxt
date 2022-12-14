package store

import (
	"main/dto"
	"main/entity"
	service "main/service"
)

type UserStoreImpl struct {
	UserService service.UserService
}

func NewUserStore(userService service.UserService) UserStore {
	return &UserStoreImpl{UserService: userService}
}

func (f *UserStoreImpl) GetAllUser() []entity.User {
	return f.UserService.GetAllUser()
}

func (u *UserStoreImpl) CreateUser(createUserDto dto.CreateUserDto) entity.User {
	return u.UserService.CreateUser(createUserDto)
}
