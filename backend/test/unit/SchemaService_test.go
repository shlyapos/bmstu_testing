package test

import (
	"skema/app/model/builder"
	"skema/app/repo"
	"skema/app/service"
	"skema/database"
	"testing"
)

func initSchemaRepos() (repo.ISchemaRepo, repo.ICommentRepo, error) {
	db, err := database.StubConnection()

	if err != nil {
		return nil, nil, err
	}

	return repo.NewSchemaRepo(db), repo.NewCommentRepo(db), nil
}

func initSchemaTestData(repo repo.ISchemaRepo) {
	objSchemaMother := new(builder.SchemaMother)

	repo.Create(objSchemaMother.Obj0())
	repo.Create(objSchemaMother.Obj1())
	repo.Create(objSchemaMother.Obj2())
	repo.Create(objSchemaMother.Obj3())
}

func TestGetSchemaById(t *testing.T) {
	schemaRepo, commentRepo, err := initSchemaRepos()

	if err != nil {
		t.Error("Error: ", err)
	}

	initSchemaTestData(schemaRepo)
	serv := service.NewSchemaService(schemaRepo, commentRepo)

	schema, err := serv.TakeSchemaById(2)

	if err != nil || schema == nil {
		t.Error("Error take user by id", schema)
	}
}

func TestGetSchemaByRating(t *testing.T) {
	schemaRepo, commentRepo, err := initSchemaRepos()

	if err != nil {
		t.Error("Error: ", err)
	}

	initSchemaTestData(schemaRepo)
	serv := service.NewSchemaService(schemaRepo, commentRepo)

	schemes, err := serv.TakeUserSchemes(1)

	if err != nil || len(*schemes) == 0 {
		t.Error("Error: with filtering by user id")
	}
}
