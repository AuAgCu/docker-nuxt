package store

import (
	"main/entity"
	service "main/service/interface"
)

type UserStoreImpl struct {
	userService service.UserService
}

func (f *UserStoreImpl) GetAllUser() []entity.User {
	return f.userService.GetAllUser()
}
