package service

import (
	"main/dto"
	"main/entity"
	repository "main/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (f *UserServiceImpl) GetAllUser() []entity.User {
	return f.UserRepository.GetAllUser()
}

func (s *UserServiceImpl) CreateUser(userDto dto.CreateUserDto) entity.User {
	u := userDto.ConvertUser()
	return s.UserRepository.CreateUser(u)
}
