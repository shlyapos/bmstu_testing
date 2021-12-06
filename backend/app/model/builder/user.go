package builder

import "describe.me/app/model"

type UserBuilder struct {
	Id     uint
	Login  string
	Name   string
	Email  string
	Rating int
}

func (b *UserBuilder) Build() *model.User {
	return &model.User{
		Id:     b.Id,
		Login:  b.Login,
		Name:   b.Name,
		Email:  b.Email,
		Rating: b.Rating,
	}
}

type UserMother struct{}

func (m *UserMother) Obj0() *model.User {
	builder := UserBuilder{
		Id:     1,
		Login:  "shlyapik",
		Name:   "Alexander Ivanov",
		Email:  "vanya.chuha@mail.ru",
		Rating: 12,
	}

	return builder.Build()
}

func (m *UserMother) Obj1() *model.User {
	builder := UserBuilder{
		Id:     2,
		Login:  "vad_stoke",
		Name:   "Vladislav Stokov",
		Email:  "vadst@gmail.ru",
		Rating: 22,
	}

	return builder.Build()
}

func (m *UserMother) Obj2() *model.User {
	builder := UserBuilder{
		Id:     3,
		Login:  "toxa_31",
		Name:   "Anton Logvinov",
		Email:  "to31xa@mail.ru",
		Rating: 9,
	}

	return builder.Build()
}
