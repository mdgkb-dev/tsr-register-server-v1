package registerquerytoregisterproperty

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"
)

type IService interface {
	CreateMany([]*models.RegisterQueryToRegisterProperty) error
}

type IRepository interface {
	createMany([]*models.RegisterQueryToRegisterProperty) error
	upsertMany([]*models.RegisterQueryToRegisterProperty) error
	deleteMany([]string) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
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
