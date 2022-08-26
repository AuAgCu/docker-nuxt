package service

import "main/entity"

type UserServiceImpl struct{}

func (f *UserServiceImpl) GetAllUser() []entity.User {
	return []entity.User{}
}
