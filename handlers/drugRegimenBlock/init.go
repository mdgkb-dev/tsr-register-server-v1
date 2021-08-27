package drugRegimenBlock

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany([]*models.DrugRegimenBlock) error
	UpsertMany([]*models.DrugRegimenBlock) error
	DeleteMany([]string) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany([]*models.DrugRegimenBlock) error
	upsertMany([]*models.DrugRegimenBlock) error
	deleteMany([]string) error
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
