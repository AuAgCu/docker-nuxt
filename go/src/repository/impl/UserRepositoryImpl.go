package repository

import (
	"main/entity"

	"github.com/jinzhu/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u *UserRepositoryImpl) GetAllUser() []entity.User {
	var users []entity.User
	u.db.Table("user").Find(&users)
	return users
}
