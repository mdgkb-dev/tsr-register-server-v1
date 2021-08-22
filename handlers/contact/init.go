package contact

import (
	"context"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IService interface {
	Create(*models.Contact) error
	Update(*models.Contact) error
	Upsert(*models.Contact) error

	CreateMany([]*models.Contact) error
	UpsertMany([]*models.Contact) error
}

type IRepository interface {
	create(*models.Contact) error
	update(*models.Contact) error
	upsert(*models.Contact) error

	createMany([]*models.Contact) error
	upsertMany([]*models.Contact) error
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
