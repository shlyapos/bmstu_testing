package service

import (
	"skema/app/model"
	"skema/app/repo"
)

type UserService struct {
	repo repo.IUserRepo
}

func NewUserService(repo repo.IUserRepo) *UserService {
	serv := UserService{repo}
	return &serv
}

func (s *UserService) CreateUser(data *model.User) error {
	return s.repo.Create(data)
}

func (s *UserService) TakeUserById(id uint) (*model.User, error) {
	user, err := s.repo.TakeOneById(id)
	return user, err
}

func (s *UserService) TakeUserRatingMoreThan(rating int) (*[]model.User, error) {
	users, err := s.repo.TakeByRatingMoreThan(rating)
	return users, err
}

// mockery -dir ./repo/ -all
