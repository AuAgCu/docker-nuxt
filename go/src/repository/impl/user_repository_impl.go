package repository

import (
	"log"
	"main/entity"

	"gorm.io/gorm"
)

const table_name = "t_user"

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u *UserRepositoryImpl) GetAllUser() []entity.User {
	var users []entity.User
	u.Db.Table(table_name).Find(&users)
	log.Println(users)
	return users
}

func (u *UserRepositoryImpl) CreateUser(user entity.User) entity.User {
	log.Println(user)
	u.Db.Table(table_name).Create(user)
	return user
}
