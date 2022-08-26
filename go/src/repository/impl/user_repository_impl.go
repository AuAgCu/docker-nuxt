package repository

import (
	"main/entity"

	"github.com/jinzhu/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u *UserRepositoryImpl) GetAllUser() []entity.User {
	var users []entity.User
	u.Db.Table("user").Find(&users)
	return users
}

func (u *UserRepositoryImpl) CreateUser(user entity.User) entity.User {
	u.Db.Table("user").Create(user)
	return user
}
