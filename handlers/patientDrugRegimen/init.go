package patientDrugRegimen

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany([]*models.PatientDrugRegimen) error
	UpsertMany([]*models.PatientDrugRegimen) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	createMany([]*models.PatientDrugRegimen) error
	upsertMany([]*models.PatientDrugRegimen) error
	deleteMany([]uuid.UUID) error
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
