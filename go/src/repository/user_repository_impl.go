package repository

import (
	"main/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{Db: DB}
}

const table_name = "t_user"

func (u *UserRepositoryImpl) GetAllUser() []entity.User {
	var users []entity.User
	u.Db.Table(table_name).Find(&users)
	return users
}

func (u *UserRepositoryImpl) CreateUser(user entity.User) entity.User {
	u.Db.Table(table_name).Create(&user)
	return user
}
