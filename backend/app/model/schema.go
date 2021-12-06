package model

type Schema struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Owner string `json:"owner" gorm:"foreignkey:Login"`
	Name  string `json:"name"`
}
