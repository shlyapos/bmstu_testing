package test

import (
	"testing"

	"skema/app/model"
	"skema/app/model/builder"
	"skema/app/repo"
	"skema/app/service"
	"skema/database"
	"skema/mocks"
	"skema/util"
)

func initStubRepo() (repo.IUserRepo, error) {
	db, err := database.StubConnection()

	if err != nil {
		return nil, err
	}

	return repo.NewUserRepo(db), nil
}

func initUserTestData(repo repo.IUserRepo) {
	objMother := new(builder.UserMother)

	repo.Create(objMother.Obj0())
	repo.Create(objMother.Obj1())
	repo.Create(objMother.Obj2())
}

// Classic style with stab

func TestGetUserById(t *testing.T) {
	repo, err := initStubRepo()

	if err != nil {
		t.Error("Error: ", err)
	}

	initUserTestData(repo)
	serv := service.NewUserService(repo)

	user, err := serv.TakeUserById(1)

	if err != nil || user == nil {
		t.Error("Error take user by id")
	}
}

func TestGetUserByRating(t *testing.T) {
	repo, err := initStubRepo()

	if err != nil {
		t.Error("Error: ", err)
	}

	initUserTestData(repo)
	serv := service.NewUserService(repo)

	users, err := serv.TakeUserRatingMoreThan(10)

	if err != nil || len(*users) == 0 {
		t.Error("Error: with filtering by user rating")
	}
}

// London style with mock

func TestCreateUserNil(t *testing.T) {
	mockRepo := new(mocks.IUserRepo)

	mockRepo.On("Create", (*model.User)(nil)).Return(util.ErrorInvalidData).Once()
	serv := service.NewUserService(mockRepo)

	err := serv.CreateUser(nil)

	if err == nil {
		t.Error("Error create with invalid data")
	}

	mockRepo.AssertExpectations(t)
}

func TestCreateUserOk(t *testing.T) {
	mockRepo := new(mocks.IUserRepo)
	objMother := new(builder.UserMother)

	mockRepo.On("Create", objMother.Obj0()).Return(nil).Once()
	serv := service.NewUserService(mockRepo)

	err := serv.CreateUser(objMother.Obj0())

	if err != nil {
		t.Error("Error database create with correct data")
	}

	mockRepo.AssertExpectations(t)
}
