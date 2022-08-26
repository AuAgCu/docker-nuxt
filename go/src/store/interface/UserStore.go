package store

import (
	"main/dto"
	"main/entity"
)

type UserStore interface {
	GetAllUser() []entity.User
	CreateUser(createUserDto dto.CreateUserDto) entity.User
}
