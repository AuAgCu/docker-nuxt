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
