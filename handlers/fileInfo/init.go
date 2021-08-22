package fileInfo

import (
	"context"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IService interface {
	Create(*models.FileInfo) error
	Update(*models.FileInfo) error
	Upsert(*models.FileInfo) error

	CreateMany([]*models.FileInfo) error
	UpsertMany([]*models.FileInfo) error
}

type IRepository interface {
	create(*models.FileInfo) error
	update(*models.FileInfo) error
	upsert(*models.FileInfo) error

	createMany([]*models.FileInfo) error
	upsertMany([]*models.FileInfo) error
}

type Handler struct {
	service IService
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
