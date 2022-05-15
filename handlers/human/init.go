package human

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IService interface {
	Create(*models.Human) error
	Update(*models.Human) error
	Delete(uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

type IHistoryRepository interface {
	getDB() *bun.DB
	create(*models.HumanHistory) error
}

type IHistoryService interface {
	Create(*models.HumanHistory) error
}

type HistoryService struct {
	repository IHistoryRepository
}

type HistoryRepository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func CreateHistoryService(db *bun.DB) *HistoryService {
	repo := NewHistoryRepository(db)
	return NewHistoryService(repo)
}

func NewHistoryService(repository IHistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func NewHistoryRepository(db *bun.DB) *HistoryRepository {
	return &HistoryRepository{db: db, ctx: context.Background()}
}
