package entity

type User struct {
	UserId    int64  `json:"userId" gorm:"primarykey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
