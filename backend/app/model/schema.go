package model

type Schema struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Owner uint   `json:"owner" gorm:"foreignkey:Login"`
	Name  string `json:"name"`

	Comments []Comment `json:"Comments"`
}
