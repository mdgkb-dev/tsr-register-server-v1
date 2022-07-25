package human

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IService interface {
	Create(*models.Human) error
	Update(*models.Human) error
	Delete(uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.Human) error
	update(*models.Human) error
	delete(uuid.UUID) error
	get(uuid.UUID) (*models.Human, error)
}

type Handler struct {
	service        IService
	historyService IHistoryService
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type IHistoryRepository interface {
	db() *bun.DB
	create(*models.HumanHistory) error
}

type IHistoryService interface {
	Create(*models.HumanHistory) error
}

type HistoryService struct {
	repository IHistoryRepository
}

type HistoryRepository struct {
	helper *helper.Helper
	ctx    context.Context
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

func CreateHistoryService() *HistoryService {
	repo := NewHistoryRepository()
	return NewHistoryService(repo)
}

func NewHistoryService(repository IHistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{ctx: context.Background()}
}
