package repository

import (
	"main/entity"
)

type UserRepository interface {
	GetAllUser() []entity.User
	CreateUser(user entity.User) entity.User
}
