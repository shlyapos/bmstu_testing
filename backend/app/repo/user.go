package repo

import (
	"skema/app/model"

	"github.com/jinzhu/gorm"
)

type IUserRepo interface {
	Create(data *model.User) error
	TakeOneById(id uint) (*model.User, error)
	TakeByRatingMoreThan(rating int) (*[]model.User, error)
}

type UserRepo struct {
	Database *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	repo := new(UserRepo)
	repo.Database = db

	return repo
}

func (repo *UserRepo) Create(data *model.User) error {
	return repo.Database.Create(data).Error
}

func (repo *UserRepo) TakeOneById(id uint) (*model.User, error) {
	user := new(model.User)
	err := repo.Database.Where("id = ?", id).First(user).Error

	return user, err
}

func (repo *UserRepo) TakeByRatingMoreThan(rating int) (*[]model.User, error) {
	users := new([]model.User)
	err := repo.Database.Where("rating >= ?", rating).Find(&users).Error

	return users, err
}
