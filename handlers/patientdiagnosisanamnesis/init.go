package patientdiagnosisanamnesis

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/uptrace/bun"
)

type IService interface {
	CreateMany([]*models.PatientDiagnosisAnamnesis) error
	UpsertMany([]*models.PatientDiagnosisAnamnesis) error
	DeleteMany([]string) error
}

type IRepository interface {
	db() *bun.DB
	createMany([]*models.PatientDiagnosisAnamnesis) error
	upsertMany([]*models.PatientDiagnosisAnamnesis) error
	deleteMany([]string) error
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