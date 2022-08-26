package store

import (
	"main/entity"
	store "main/store/interface"
)

type UserStoreImpl struct {
	UserStore store.UserStore
}

func (f *UserStoreImpl) GetAllUser() []entity.User {
	return f.UserStore.GetAllUser()
}
