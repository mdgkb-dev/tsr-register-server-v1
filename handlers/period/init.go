package period

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(*models.Period) error
	Update(*models.Period) error

	CreateMany([]*models.Period) error
	UpsertMany([]*models.Period) error
}

type IRepository interface {
	create(*models.Period) error
	update(*models.Period) error

	createMany([]*models.Period) error
	upsertMany([]*models.Period) error
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
