package meta

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models/schema"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetCount(c *gin.Context) error
	GetSchema(c *gin.Context) error
}

type IService interface {
	GetCount(*string) (*int, error)
	GetSchema() schema.Schema
}

type IRepository interface {
	getCount(*string) (*int, error)
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
