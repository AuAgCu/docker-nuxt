package store

import (
	"main/entity"
	store "main/store/interface"
)

type UserStoreImpl struct {
	userStore store.UserStore
}

func (f *UserStoreImpl) GetAllUser() []entity.User {
	return f.userStore.GetAllUser()
}
