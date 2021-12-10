package model

type Comment struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Owner  uint   `json:"owner" gorm:"foreignkey:User"`
	Schema uint   `json:"schema" gorm:"foreignkey:Schema"`
	Data   string `json:"data"`
}
