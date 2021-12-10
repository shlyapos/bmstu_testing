package repo

import (
	"skema/app/model"

	"github.com/jinzhu/gorm"
)

type ICommentRepo interface {
	Create(data *model.Comment) error
	TakeByUserId(id uint) (*[]model.Comment, error)
	FilterBySchemaAndUserId(schemaId uint, owner uint) (*[]model.Comment, error)
}

type CommentRepo struct {
	Database *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	repo := new(CommentRepo)
	repo.Database = db

	return repo
}

func (repo *CommentRepo) Create(data *model.Comment) error {
	return repo.Database.Create(data).Error
}

func (repo *CommentRepo) TakeByUserId(id uint) (*[]model.Comment, error) {
	comments := new([]model.Comment)
	err := repo.Database.Where("owner = ?", id).Find(&comments).Error

	return comments, err
}

func (repo *CommentRepo) FilterBySchemaAndUserId(schemaId uint, owner uint) (*[]model.Comment, error) {
	comments := new([]model.Comment)
	err := repo.Database.Where("owner = ? AND schema = ?", owner, schemaId).Find(&comments).Error

	return comments, err
}
