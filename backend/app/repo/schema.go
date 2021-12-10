package repo

import (
	"skema/app/model"

	"github.com/jinzhu/gorm"
)

type ISchemaRepo interface {
	Create(data *model.Schema) error
	TakeOneById(id uint) (*model.Schema, error)
	FilterByOwner(ownerId uint) (*[]model.Schema, error)
}

type SchemaRepo struct {
	Database *gorm.DB
}

func NewSchemaRepo(db *gorm.DB) *SchemaRepo {
	repo := new(SchemaRepo)
	repo.Database = db

	return repo
}

func (repo *SchemaRepo) Create(data *model.Schema) error {
	return repo.Database.Create(data).Error
}

func (repo *SchemaRepo) TakeOneById(id uint) (*model.Schema, error) {
	schema := new(model.Schema)
	err := repo.Database.Where("id = ?", id).First(schema).Error

	return schema, err
}

func (repo *SchemaRepo) FilterByOwner(ownerId uint) (*[]model.Schema, error) {
	schemes := new([]model.Schema)
	err := repo.Database.Where("owner = ?", ownerId).Find(&schemes).Error

	return schemes, err
}
