package service

import (
	"main/entity"
	repository "main/repository/interface"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (f *UserServiceImpl) GetAllUser() []entity.User {
	return f.userRepository.GetAllUser()
}
