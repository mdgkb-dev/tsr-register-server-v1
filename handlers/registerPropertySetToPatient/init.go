package registerPropertySetToPatient

import (
	"context"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IService interface {
	CreateMany([]*models.RegisterPropertySetToPatient) error
}

type IRepository interface {
	getDB() *bun.DB
	createMany([]*models.RegisterPropertySetToPatient) error
	upsertMany([]*models.RegisterPropertySetToPatient) error
	deleteMany([]string) error
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