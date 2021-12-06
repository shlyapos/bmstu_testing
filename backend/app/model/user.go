package model

type User struct {
	Id     uint   `json:"id" gorm:"primary_key"`
	Login  string `json:"login" gorm:"unique"`
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"unique"`
	Rating int    `json:"rating" gorm:"default:0"`

	Schemes []Schema `json:"Schemes"`
}

// type User struct {
// 	id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
// 	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
// 	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
// 	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
// }

// func NewUser(firstName string, lastName string, userName string, email string) *User {
// 	return &User{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Username:  userName,
// 		Email:     email,
// 	}
// }
