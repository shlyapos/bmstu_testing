package test

import (
	"skema/app/model/builder"
	"skema/app/repo"
	"skema/app/service"
	"skema/database"
	"testing"
)

func initDataSchema(repo repo.ISchemaRepo) {
	objMother := new(builder.SchemaMother)

	repo.Create(objMother.Obj0())
	repo.Create(objMother.Obj1())
	repo.Create(objMother.Obj2())
	repo.Create(objMother.Obj3())
}

func initDataUser(repo repo.IUserRepo) {
	objMother := new(builder.UserMother)

	repo.Create(objMother.Obj0())
	repo.Create(objMother.Obj1())
	repo.Create(objMother.Obj2())
}

func initDataComment(repo repo.ICommentRepo) {
	objMother := new(builder.CommentMother)

	repo.Create(objMother.Obj0())
	repo.Create(objMother.Obj1())
	repo.Create(objMother.Obj2())
	repo.Create(objMother.Obj3())
	repo.Create(objMother.Obj4())
	repo.Create(objMother.Obj5())
}

func initStubRepos() (repo.IUserRepo, repo.ISchemaRepo, repo.ICommentRepo) {
	db, err := database.StubConnection()

	if err != nil {
		panic(err)
	}

	userRepo := repo.NewUserRepo(db)
	schemaRepo := repo.NewSchemaRepo(db)
	commentRepo := repo.NewCommentRepo(db)

	return userRepo, schemaRepo, commentRepo
}

func TestUserProfileByLoginSuccess(t *testing.T) {
	userRepo, schemaRepo, _ := initStubRepos()

	initDataUser(userRepo)
	initDataSchema(schemaRepo)

	userService := service.NewUserService(userRepo)
	schemaService := service.NewSchemaService(schemaRepo, nil)

	user, userErr := userService.TakeUserById(2)
	schemes, schemaErr := schemaService.TakeUserSchemes(user.Id)

	if userErr != nil || user == nil {
		t.Error("Error user: ", userErr, user)
	}

	if schemaErr != nil || schemes == nil || len(*schemes) == 0 {
		t.Error("Error schemes: ", schemaErr, schemes)
	}
}

func TestUserCommentStatisticSuccess(t *testing.T) {
	userRepo, _, commentRepo := initStubRepos()

	initDataUser(userRepo)
	initDataComment(commentRepo)

	userService := service.NewUserService(userRepo)
	schemaService := service.NewSchemaService(nil, commentRepo)

	user, userErr := userService.TakeUserById(3)
	comments, commentErr := schemaService.TakeCommentsByUserId(user.Id)

	if userErr != nil || user == nil {
		t.Error("Error user: ", userErr, user)
	}

	if commentErr != nil || comments == nil || len(*comments) == 0 {
		t.Error("Error schemes: ", commentErr, comments)
	}
}

func TestUserGlobalStatisticSuccess(t *testing.T) {
	userRepo, schemaRepo, commentRepo := initStubRepos()

	initDataUser(userRepo)
	initDataSchema(schemaRepo)
	initDataComment(commentRepo)

	userService := service.NewUserService(userRepo)
	schemaService := service.NewSchemaService(schemaRepo, commentRepo)

	user, userErr := userService.TakeUserById(3)
	schemes, schemaErr := schemaService.TakeUserSchemes(user.Id)
	comments, commentErr := schemaService.TakeCommentsByUserId(user.Id)

	if userErr != nil || user == nil {
		t.Error("Error user: ", userErr, user)
	}

	if schemaErr != nil || schemes == nil || len(*schemes) == 0 {
		t.Error("Error schemes: ", schemaErr, schemes)
	}

	if commentErr != nil || comments == nil || len(*comments) == 0 {
		t.Error("Error schemes: ", commentErr, comments)
	}
}

func TestFilterCommentForCurrentUserAndSchema(t *testing.T) {
	userRepo, schemaRepo, commentRepo := initStubRepos()

	initDataUser(userRepo)
	initDataSchema(schemaRepo)
	initDataComment(commentRepo)

	userService := service.NewUserService(userRepo)
	schemaService := service.NewSchemaService(schemaRepo, commentRepo)

	user, userErr := userService.TakeUserById(3)
	schema, schemaErr := schemaService.TakeSchemaById(2)
	comments, commentErr := schemaService.TakeSchemaCommentByUserId(schema.Id, user.Id)

	if userErr != nil || user == nil {
		t.Error("Error user: ", userErr, user)
	}

	if schemaErr != nil || schema == nil {
		t.Error("Error schemes: ", schemaErr, schema)
	}

	if commentErr != nil || comments == nil || len(*comments) == 0 {
		t.Error("Error schemes: ", commentErr, comments)
	}
}
