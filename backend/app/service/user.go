package service

import (
	"skema/app/model"
	"skema/app/repo"
	"skema/util"
)

type UserService struct {
	repo repo.IUserRepo
}

// var infoUserLogger = util.NewInfoLogger()

func NewUserService(repo repo.IUserRepo) *UserService {
	serv := UserService{repo}
	return &serv
}

func (s *UserService) CreateUser(data *model.User) error {
	err := s.repo.Create(data)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "UserService", "CreateUser")
	return err
}

func (s *UserService) TakeUserById(id uint) (*model.User, error) {
	user, err := s.repo.TakeOneById(id)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "UserService", "TakeUserById")
	return user, err
}

func (s *UserService) TakeUserRatingMoreThan(rating int) (*[]model.User, error) {
	users, err := s.repo.TakeByRatingMoreThan(rating)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "UserService", "TakeUserRatingMoreThan")
	return users, err
}
