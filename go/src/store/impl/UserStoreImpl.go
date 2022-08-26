package store

import "main/entity"

type UserStoreImpl struct{}

func (f *UserStoreImpl) GetAllUser() []entity.User {
	return []entity.User{}
}
