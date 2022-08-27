package service

import (
	"main/dto"
	"main/entity"
)

type UserService interface {
	GetAllUser() []entity.User
	CreateUser(userDto dto.CreateUserDto) entity.User
}
