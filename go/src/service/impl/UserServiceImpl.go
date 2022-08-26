package service

import (
	"main/entity"
	repository "main/repository/interface"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (f *UserServiceImpl) GetAllUser() []entity.User {
	return f.UserRepository.GetAllUser()
}
