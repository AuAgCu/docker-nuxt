package repository

import "main/entity"

type UserRepository interface {
	GetAllUser() []entity.User
}
