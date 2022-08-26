package dto

import "main/entity"

type CreateUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (v CreateUserDto) ConvertUser() entity.User {
	return entity.User{
		FirstName: v.FirstName,
		LastName:  v.LastName,
	}
}
