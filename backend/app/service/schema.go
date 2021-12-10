package service

import (
	"skema/app/model"
	"skema/app/repo"
	"skema/util"
)

type SchemaService struct {
	schemaRepo  repo.ISchemaRepo
	commentRepo repo.ICommentRepo
}

func NewSchemaService(schemaRepo repo.ISchemaRepo, commentRepo repo.ICommentRepo) *SchemaService {
	serv := SchemaService{schemaRepo, commentRepo}
	return &serv
}

func (s *SchemaService) CreateSchema(data *model.Schema) error {
	err := s.schemaRepo.Create(data)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "SchemaService", "CreateSchema")
	return err
}

func (s *SchemaService) TakeSchemaById(id uint) (*model.Schema, error) {
	schema, err := s.schemaRepo.TakeOneById(id)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "SchemaService", "TakeSchemaById")
	return schema, err
}

func (s *SchemaService) TakeUserSchemes(owner uint) (*[]model.Schema, error) {
	schemes, err := s.schemaRepo.FilterByOwner(owner)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "SchemaService", "TakeUserSchemes")
	return schemes, err
}

func (s *SchemaService) TakeCommentsByUserId(id uint) (*[]model.Comment, error) {
	comments, err := s.commentRepo.TakeByUserId(id)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "SchemaService", "TakeCommentsByUserId")
	return comments, err
}

func (s *SchemaService) TakeSchemaCommentByUserId(schemaId uint, userId uint) (*[]model.Comment, error) {
	comments, err := s.commentRepo.FilterBySchemaAndUserId(schemaId, userId)
	status := "Success"

	if err != nil {
		status = "Error"
	}

	util.Logger.WriteInfoLog(status, "SchemaService", "TakeSchemaCommentByUserId")
	return comments, err
}
