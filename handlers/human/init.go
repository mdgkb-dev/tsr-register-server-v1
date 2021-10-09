package human

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

type IService interface {
	Create(*models.Human) error
	Update(*models.Human) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Human) error
	update(*models.Human) error
}

type Handler struct {
	service        IService
	historyService IHistoryService
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
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
