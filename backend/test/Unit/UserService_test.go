package test

import (
	"testing"

	"describe.me/app/model"
	"describe.me/app/model/builder"
	"describe.me/app/repo"
	"describe.me/app/service"
	"describe.me/database"
	"describe.me/mocks"
	"describe.me/util"
)

func InitStabRepo() (repo.IUserRepo, error) {
	db, err := database.StubConnection()

	if err != nil {
		return nil, err
	}

	return repo.NewUserRepo(db), nil
}

func InitMockRepo() repo.IUserRepo {
	return new(mocks.IUserRepo)
}

func InitTestData(serv *service.UserService) {
	objMother := new(builder.UserMother)

	serv.CreateUser(objMother.Obj0())
	serv.CreateUser(objMother.Obj1())
	serv.CreateUser(objMother.Obj2())
}

// Classic style with stab

func TestGetUserById(t *testing.T) {
	repo, err := InitStabRepo()

	if err != nil {
		t.Error("Error: ", err)
	}

	serv := service.NewUserService(repo)
	InitTestData(serv)

	user, err := serv.TakeUserById(1)

	if err != nil || user == nil {
		t.Error("Error take user by id")
	}
}

func TestGetUserByRating(t *testing.T) {
	repo, err := InitStabRepo()

	if err != nil {
		t.Error("Error: ", err)
	}

	serv := service.NewUserService(repo)
	InitTestData(serv)

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
