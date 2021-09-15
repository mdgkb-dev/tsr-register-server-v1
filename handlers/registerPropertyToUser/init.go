package registerPropertyToUser

import (
	"context"
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

type IHandler interface {
	Create(*gin.Context) error
	Delete(*gin.Context) error
}

type IService interface {
	Create(*models.RegisterPropertyToUser) error
	Delete(*models.RegisterPropertyToUser) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterPropertyToUser) error
	delete(*models.RegisterPropertyToUser) error
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

func CreateHandler(db *bun.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service)
}

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
	return &Handler{service: s}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
