package store

import "main/entity"

type UserStore interface {
	GetAllUser() []entity.User
}
