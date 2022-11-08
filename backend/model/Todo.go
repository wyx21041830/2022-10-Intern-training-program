package model

type User struct {
	UserId   uint   `json:"id" form:"id" gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Password string `json:"passwd" form:"passwd"`
}

type Todo struct {
	Id     uint   `json:"id" form:"id" gorm:"primaryKey"`
	UserId uint   `json:"userId" form:"userId"`
	Info   string `json:"Info" form:"Info"`
}