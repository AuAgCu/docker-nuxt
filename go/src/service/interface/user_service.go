package service

import "main/entity"

type UserService interface {
	GetAllUser() []entity.User
}
