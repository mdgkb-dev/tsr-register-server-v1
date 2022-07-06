package contact

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

type IService interface {
	Create(*models.Contact) error
	Update(*models.Contact) error
	Upsert(*models.Contact) error
	Delete(uuid.NullUUID) error

	CreateMany([]*models.Contact) error
	UpsertMany([]*models.Contact) error
}

type IRepository interface {
	create(*models.Contact) error
	update(*models.Contact) error
	upsert(*models.Contact) error
	delete(uuid.NullUUID) error

	createMany([]*models.Contact) error
	upsertMany([]*models.Contact) error
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
