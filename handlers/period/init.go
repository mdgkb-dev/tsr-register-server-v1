package period

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"
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
