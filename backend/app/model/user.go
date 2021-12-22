package model

type User struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Login  string `json:"login" gorm:"unique"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Rating int    `json:"rating" gorm:"default:0"`

	Schemes []Schema `json:"Schemes"`
}
